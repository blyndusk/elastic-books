package server

import (
	"fmt"
	"net/http"

	"github.com/blyndusk/elastic-books/api/es"
	"github.com/blyndusk/elastic-books/api/models"
	"github.com/gin-gonic/gin"
)

func SearchBook(c *gin.Context) {
	// get params
	query := c.Query("query")
	searchType := c.Query("type")

	// get result
	foundBooks := es.SearchBook(query, searchType)

	// handle no result
	if foundBooks != nil {
		c.JSON(http.StatusOK, gin.H{
			"_message": "Here is your search results",
			"data":     foundBooks,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"_message": "No result found. Try differents queries or types.",
		})
	}
}

func CreateBook(c *gin.Context) {
	// get params
	bookToCreate := models.Book{
		Id:     "",
		Name:   c.Query("name"),
		Author: c.Query("author"),
		Resume: c.Query("resume"),
	}

	// get created book from es
	createdBook := es.CreateBook(bookToCreate)

	c.JSON(http.StatusOK, gin.H{
		"_message": fmt.Sprintf("Book created [%s]", createdBook.Id),
		"data":     createdBook,
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
	// get params
	bookToUpdate := models.Book{
		Id:     c.Params.ByName("id"),
		Name:   c.Query("name"),
		Author: c.Query("author"),
		Resume: c.Query("resume"),
	}

	// get updated book from es
	updatedBook, err := es.UpdateBook(bookToUpdate)

	// handle inexisting book
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"_message": fmt.Sprintf("Book [%s] doesn't exist. Please use a valid ID.", updatedBook.Id),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"_message": fmt.Sprintf("Book updated [%s]", updatedBook.Id),
			"data":     updatedBook,
		})
	}
}

func DeleteBook(c *gin.Context) {
	id := c.Params.ByName("id")
	response := es.DeleteBook(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted",
		"data":    response,
	})
}
