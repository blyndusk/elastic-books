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

func Search() {
	// init es client
	ctx := context.Background()
	esclient, err := InitClient()
	helpers.ExitOnError("client init", err)

	// creating book object
	newBook := models.Book{
		Name:   "Martine veut dissoudre l'Assemblée Nationale, feat Jean Castête ",
		Author: "Jeanne Oskour",
		Resume: "C'est l'histoire de Martine qui veut dissoudre l'Assemblée Nationale, accompagnée de son fidèle écuyer, Jean Castête. Un fabuleux livre pour petits et grands, de Jeanne Oskhour",
	}
	dataJSON, err := json.Marshal(newBook)
	helpers.ExitOnError("create new book", err)
	js := string(dataJSON)

	// insert new book
	_, err = esclient.Index().
		Index("books").
		BodyJson(js).
		Do(ctx)
	helpers.ExitOnError("insert new book", err)
	logrus.Info("New book inserted !")

	var books models.Books

	// init search source
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("name", "Jean"))

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