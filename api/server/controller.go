package server

import (
	"net/http"

	"github.com/blyndusk/elastic-books/api/es"
	"github.com/gin-gonic/gin"
)

func SearchBook(c *gin.Context) {

	query := c.Query("query")
	searchType := c.Query("type")
	result := es.Search(query, searchType)

	c.JSON(http.StatusOK, gin.H{
		"message": es.Esclient.String(),
		"result":  result,
	})
}
