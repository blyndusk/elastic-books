package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/blyndusk/elastic-books/api/helpers"
	elastic "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Resume string `json:"resume"`
}

func GetESClient() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL("http://es01:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err

}

func main() {

	ctx := context.Background()
	esclient, err := GetESClient()
	logrus.Info(ctx, esclient)
	helpers.ExitOnError("client init", err)

	//creating book object
	newBook := Book{
		Name:   "Martine veut dissoudre l'Assemblée Nationale, feat Jean Castête ",
		Author: "Jeanne Oskour",
		Resume: "C'est l'histoire de Martine qui veut dissoudre l'Assemblée Nationale, accompagnée de son fidèle écuyer, Jean Castête. Un fabuleux livre pour petits et grands, de Jeanne Oskhour",
	}

	dataJSON, err := json.Marshal(newBook)
	helpers.ExitOnError("create new book", err)

	js := string(dataJSON)
	ind, err := esclient.Index().
		Index("books").
		BodyJson(js).
		Do(ctx)

	helpers.ExitOnError("insert new book", err)

	logrus.Info("[Elastic][InsertProduct]Insertion Successful", ind)

	var books []Book
	logrus.Info(books)
}
