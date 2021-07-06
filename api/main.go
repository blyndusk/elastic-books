package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
)

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Resume string `json:"resume"`
}

func GetESClient() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err

}

func main() {

	ctx := context.Background()
	esclient, err := GetESClient()
	logrus.Info(ctx, esclient)
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}


}
