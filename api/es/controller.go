package es

import (
	"context"
	"encoding/json"

	"github.com/blyndusk/elastic-books/api/helpers"
	"github.com/blyndusk/elastic-books/api/models"
	elastic "github.com/olivere/elastic/v7"
)

var ctx = context.Background()

func SearchBook(query string, searchType string) models.Books {
	var foundBooks models.Books

	// init search source with query
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery(searchType, query))

	// init search service
	searchService := Esclient.Search().Index("books").SearchSource(searchSource)

	// get result with context
	searchResult, err := searchService.Do(ctx)
	helpers.ExitOnError("get search query - ", err)

	// parsing result
	for _, hit := range searchResult.Hits.Hits {
		var foundBook models.Book
		// fill result JSON
		err := json.Unmarshal(hit.Source, &foundBook)
		helpers.ExitOnError("stringify source - ", err)
		foundBook.Id = hit.Id
		foundBooks = append(foundBooks, foundBook)
	}

	return foundBooks
}

func CreateBook(bookToCreate models.Book) models.Book {
	// extract data
	data, err := json.Marshal(bookToCreate)
	helpers.ExitOnError("parsing book - ", err)
	js := string(data)

	// insert new book in index
	resp, err := Esclient.Index().
		Index("books").
		BodyJson(js).
		Do(ctx)
	helpers.ExitOnError("insert new book - ", err)

	createdBook := bookToCreate
	createdBook.Id = resp.Id

	return createdBook
}

func ReadBook(bookToRead models.Book) (models.Book, error) {
	resp, err := Esclient.Get().
		Index("books").
		Id(bookToRead.Id).
		Do(ctx)

	if err != nil {
		// return not found in server/controller.go
		return bookToRead, err
	} else {
		readedBook := bookToRead
		err = json.Unmarshal(resp.Source, &readedBook)
		helpers.ExitOnError("stringify source - ", err)

		return readedBook, err
	}
}

func UpdateBook(bookToUpdate models.Book) (models.Book, error) {
	doc := map[string]interface{}{
		"name":   bookToUpdate.Name,
		"author": bookToUpdate.Author,
		"resume": bookToUpdate.Resume,
	}

	// Update book with specified ID and doc
	resp, err := Esclient.Update().
		Index("books").
		Id(bookToUpdate.Id).
		Doc(doc).
		DetectNoop(true).
		Do(ctx)

	if err != nil {
		return bookToUpdate, err
	} else {
		updatedBook := bookToUpdate
		updatedBook.Id = resp.Id

		return updatedBook, err
	}
}

func DeleteBook(id string) error {
	_, err := Esclient.Delete().
		Index("books").
		Id(id).
		Do(ctx)

	return err
}
