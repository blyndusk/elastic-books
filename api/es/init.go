package es

import (
	"context"

	"github.com/blyndusk/elastic-books/api/helpers"
	elastic "github.com/olivere/elastic/v7"
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
	helpers.ExitOnError("set new client - ", err)

	ctx := context.Background()

	exists, err := esclient.IndexExists("books").Do(ctx)
	helpers.ExitOnError("index exist - ", err)

	if !exists {
		_, err := esclient.CreateIndex("books").BodyString(mapping).Do(ctx)
		helpers.ExitOnError("create index - ", err)
	}
	Esclient = esclient
}
