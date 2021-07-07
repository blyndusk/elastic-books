package es

import (
	"context"
	elastic "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

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
		  		"type": "integer"
			},
			"resume": {
		  		"type": "float"
			}
	  	}
	}
}`

func InitClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://es01:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	logrus.Info("ES initialized !")

	ctx := context.Background()

	// Check if "books" index exists
	exists, err := client.IndexExists("books").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex("books").BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			logrus.Info("Something went wrong :/ The \"books\" index wasn't created.")
		}
	}

	return client, err
}
