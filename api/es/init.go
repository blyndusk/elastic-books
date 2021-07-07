package es

import (
	"context"

	"github.com/blyndusk/elastic-books/api/helpers"
	elastic "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

var Esclient *elastic.Client

const mapping = `{
	"settings": {
	  	"number_of_shards": 1,
	  	"number_of_replicas": 1
	},
	"mappings": {
	  	"properties": {
			"name": {
		 		"type": "text"
			},
			"author": {
		  		"type": "text"
			},
			"resume": {
		  		"type": "text"
			}
	  	}
	}
}`

func InitClient() {
	esclient, err := elastic.NewClient(elastic.SetURL("http://es01:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	helpers.ExitOnError("new client", err)
	logrus.Info("ES initialized !")

	ctx := context.Background()

	// Check if "books" index exists
	exists, err := esclient.IndexExists("books").Do(ctx)
	helpers.ExitOnError("index exist", err)

	if !exists {
		// Create a new index.
		createIndex, err := esclient.CreateIndex("books").BodyString(mapping).Do(ctx)
		helpers.ExitOnError("create index", err)

		if !createIndex.Acknowledged {
			logrus.Info("Something went wrong :/ The \"books\" index wasn't created.")
		}
	}
	Esclient = esclient

}
