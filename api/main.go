package main

import (
	"github.com/blyndusk/elastic-books/api/es"
	"github.com/blyndusk/elastic-books/api/router"
)

func main() {
	es.InitClient()
	router.Setup()
}
