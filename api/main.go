package main

import (
	"github.com/blyndusk/elastic-books/api/es"
	"github.com/blyndusk/elastic-books/api/helpers"
)

func main() {
	// init es client
	_, err := es.InitClient()
	helpers.ExitOnError("client init", err)
}
