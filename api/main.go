package main

import (
	"github.com/blyndusk/elastic-books/api/es"
	"github.com/blyndusk/elastic-books/api/router"
	"github.com/sirupsen/logrus"
)

func main() {
	// init elastic search client
	es.InitClient()
	// setup router
	router.Setup()

	logrus.Info("Thanks for using elastic-books !")
}
