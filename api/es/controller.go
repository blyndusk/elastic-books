package es

import (
	"context"
	"encoding/json"

	"github.com/blyndusk/elastic-books/api/helpers"
	"github.com/blyndusk/elastic-books/api/models"
	elastic "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

var Ctx = context.Background()

func SearchBook(query string, searchType string) models.Books {
	var foundBooks models.Books

	// init search source with query
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery(searchType, query))

	// init search service
	searchService := Esclient.Search().Index("books").SearchSource(searchSource)

	// get result with context
	searchResult, err := searchService.Do(Ctx)
	helpers.ExitOnError("get search query", err)

	// parsing result
	for _, hit := range searchResult.Hits.Hits {
		var book models.Book
		// fill result JSON
		err := json.Unmarshal(hit.Source, &book)
		helpers.ExitOnError("stringify source", err)
		foundBooks = append(foundBooks, book)
	}

	return foundBooks
}

func CreateBook(bookToCreate models.Book) models.Book {
	// extract data
	data, err := json.Marshal(bookToCreate)
	helpers.ExitOnError("parsing book ", err)
	js := string(data)

	// insert new book in index
	resp, err := Esclient.Index().
		Index("books").
		BodyJson(js).
		Do(Ctx)
	helpers.ExitOnError("insert new book", err)

	// get created book + add id
	createdBook := bookToCreate
	createdBook.Id = resp.Id

	return createdBook
}

func ReadBook(id string) *elastic.GetResult {
	// Read book with specified ID
	book, err := Esclient.Get().
		Index("books").
		Id(id).
		Do(Ctx)

	helpers.ExitOnError("Read Book", err)
	if book.Found {
		logrus.Info("Book found: \n %s", book.Fields)
	}
	return book
}

func UpdateBook(bookToUpdate models.Book) (models.Book, error) {
	// extract data
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
		Do(Ctx)

	// if err (not found), return with book param + err
	if err != nil {
		return bookToUpdate, err
	} else {
		// get updated book + add id
		updatedBook := bookToUpdate
		updatedBook.Id = resp.Id

		return updatedBook, err
	}
}

func DeleteBook(id string) *elastic.DeleteResponse {
	// Delete book with specified ID
	resp, err := Esclient.Delete().
		Index("books").
		Id(id).
		Do(Ctx)
	helpers.ExitOnError("Delete Book", err)
	logrus.Info("Book has been deleted !")
	return resp
}
