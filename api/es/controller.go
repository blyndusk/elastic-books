package es

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/blyndusk/elastic-books/api/helpers"
	"github.com/blyndusk/elastic-books/api/models"
	elastic "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

var Ctx = context.Background()

func SearchBook(query string, searchType string) models.Books {
	var books models.Books

	// init search source
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery(searchType, query))

	// search query
	queryStr, err1 := searchSource.Source()
	queryJs, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		logrus.Fatal("query", err1, err2)
	}
	logrus.Info("Final ESQuery ", string(queryJs))

	// init search service
	searchService := Esclient.Search().Index("books").SearchSource(searchSource)

	searchResult, err := searchService.Do(Ctx)
	helpers.ExitOnError("get search query", err)

	// searching
	for _, hit := range searchResult.Hits.Hits {
		var book models.Book
		err := json.Unmarshal(hit.Source, &book)
		helpers.ExitOnError("getting books", err)
		books = append(books, book)
	}
	helpers.ExitOnError("Fetching book fail", err)
	logrus.Info(books)
	for _, s := range books {
		logrus.Info(fmt.Sprintf("Book found: \nName: %s\nAuthor: %s\nResume: %s \n", s.Name, s.Author, s.Resume))
	}
	return books
}

func CreateBook(name string, author string, resume string) models.Book {
	// creating book object
	book := models.Book{
		Id:     "",
		Name:   name,
		Author: author,
		Resume: resume,
	}
	dataJSON, err := json.Marshal(book)
	helpers.ExitOnError("create new book", err)
	js := string(dataJSON)

	// insert new book
	resp, err := Esclient.Index().
		Index("books").
		BodyJson(js).
		Do(Ctx)
	book.Id = resp.Id
	helpers.ExitOnError("insert new book", err)
	logrus.Info("New book inserted !")
	return book
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

func UpdateBook(id string, name string, author string, resume string) *elastic.UpdateResponse {
	// Update book with specified ID
	resp, err := Esclient.Update().
		Index("books").
		Id(id).
		Doc(map[string]interface{}{"name": name, "author": author, "resume": resume}).
		DetectNoop(true).
		Do(Ctx)

	helpers.ExitOnError("Update Book", err)
	logrus.Info("Book has been updated !")
	return resp
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
