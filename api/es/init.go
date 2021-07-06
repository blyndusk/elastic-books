package es

import (
	elastic "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func InitClient() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL("http://es01:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	logrus.Info("ES initialized !")

	return client, err

}
