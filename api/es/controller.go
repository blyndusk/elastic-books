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

func Create(elastic.Client esclient, string Name, string Author, string Resume) {
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
	ctx := context.Background()
	_, err = esclient.Index().
		Index("books").
		BodyJson(js).
		Do(ctx)
	helpers.ExitOnError("insert new book", err)
	logrus.Info("New book inserted !")
}

func Read(elastic.Client esclient, string Query, string SearchType) {
	var books models.Books

	// init search source
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery(SearchType, Query))

	// search query
	queryStr, err1 := searchSource.Source()
	queryJs, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		logrus.Fatal("query", err, err)
	}
	logrus.Info("Final ESQuery ", string(queryJs))

	// init search service
	searchService := esclient.Search().Index("books").SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	helpers.ExitOnError("get search query", err)

	// searching
	for _, hit := range searchResult.Hits.Hits {
		var book models.Book
		err := json.Unmarshal(hit.Source, &book)
		helpers.ExitOnError("getting books", err)
		books = append(books, book)
	}
	helpers.ExitOnError("Fetching book fail", err)

	for _, s := range books {
		logrus.Info(fmt.Sprintf("Book found: \nName: %s\nAuthor: %s\nResume: %s \n", s.Name, s.Author, s.Resume))
	}
}

func Update(elastic.Client esclient, int Id, string Name, string Author, string Resume) {
	// Update book with specified ID
	newBook := models.Book{
		Name:   Name,
		Author: Author,
		Resume: Resume,
	}
	dataJSON, err := json.Marshal(newBook)
	helpers.ExitOnError("create new book", err)
	js := string(dataJSON)

	ctx := context.Background()
	_, err := client.Update().
		Index("books").
		Id(Id).
		BodyJson(js).
		Do(ctx)
	helpers.ExitOnError("Update Book", err)
	logrus.Info("Book has been updated !")
}

func Delete(elastic.Client esclient, int Id) {
	// Delete book with specified ID
	ctx := context.Background()
	_, err := esclient.Delete().
		Index("books").
		Id(Id).
		BodyJson(js).
		Do(ctx)
	helpers.ExitOnError("Delete Book", err)
	logrus.Info("Book has been deleted !")
}

