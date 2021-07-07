package main

import (
	"github.com/blyndusk/elastic-books/api/es"
)

func main() {
	// init es client
	ctx := context.Background()
	esclient, err := es.InitClient()
	helpers.ExitOnError("client init", err)
}
