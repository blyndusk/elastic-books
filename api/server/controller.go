package server

import (
	"net/http"

	"github.com/blyndusk/elastic-books/api/es"
	"github.com/gin-gonic/gin"
)

func SearchBook(c *gin.Context) {
	query := c.Query("query")
	searchType := c.Query("type")
	data := es.SearchBook(query, searchType)

	c.JSON(http.StatusOK, gin.H{
		"message": "Here is your search results",
		"data":    data,
	})
}

func CreateBook(c *gin.Context) {
	name := c.Query("name")
	author := c.Query("author")
	resume := c.Query("resume")
	data := es.CreateBook(name, author, resume)

	c.JSON(http.StatusOK, gin.H{
		"message": "New book created",
		"data":    data,
	})
}

func ReadBook(c *gin.Context) {
	id := c.Params.ByName("id")
	response := es.ReadBook(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Book",
		"data":    response,
	})
}

func UpdateBook(c *gin.Context) {
	id := c.Params.ByName("id")
	name := c.Query("name")
	author := c.Query("author")
	resume := c.Query("resume")
	response := es.UpdateBook(id, name, author, resume)

	c.JSON(http.StatusOK, gin.H{
		"message": "Book updated",
		"data":    response,
	})
}

func DeleteBook(c *gin.Context) {
	id := c.Params.ByName("id")
	response := es.DeleteBook(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted",
		"data":    response,
	})
}
