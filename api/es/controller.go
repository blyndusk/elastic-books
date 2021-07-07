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

func Create(Name string, Author string, Resume string) {
	// creating book object
	newBook := models.Book{
		Name:   Name,
		Author: Author,
		Resume: Resume,
	}
	dataJSON, err := json.Marshal(newBook)
	helpers.ExitOnError("create new book", err)
	js := string(dataJSON)

	// insert new book
	_, err = Esclient.Index().
		Index("books").
		BodyJson(js).
		Do(Ctx)
	helpers.ExitOnError("insert new book", err)
	logrus.Info("New book inserted !")
}

func Search(Query string, SearchType string) models.Books {
	var books models.Books

	// init search source
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery(SearchType, Query))

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

func Update(Id string, Name string, Author string, Resume string) {
	// Update book with specified ID
	_, err := Esclient.Update().
		Index("books").
		Id(Id).
		Doc(map[string]interface{}{"name": Name, "author": Author, "resume": Resume}).
		DetectNoop(true).
		Do(Ctx)

	helpers.ExitOnError("Update Book", err)
	logrus.Info("Book has been updated !")
}

func Delete(Id string) {
	// Delete book with specified ID
	_, err := Esclient.Delete().
		Index("books").
		Id(Id).
		Do(Ctx)
	helpers.ExitOnError("Delete Book", err)
	logrus.Info("Book has been deleted !")
}
