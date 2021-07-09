package server

import (
	"fmt"
	"net/http"

	"github.com/blyndusk/elastic-books/api/es"
	"github.com/blyndusk/elastic-books/api/models"
	"github.com/gin-gonic/gin"
)

func SearchBook(c *gin.Context) {
	// get request params
	query := c.Query("query")
	searchType := c.Query("type")

	// get result
	foundBooks := es.SearchBook(query, searchType)

	// handle no result from the search
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
	bookToCreate := models.Book{
		Id:     "",
		Name:   c.Query("name"),
		Author: c.Query("author"),
		Resume: c.Query("resume"),
	}

	createdBook := es.CreateBook(bookToCreate)

	c.JSON(http.StatusOK, gin.H{
		"_message": fmt.Sprintf("Book created [%s]", createdBook.Id),
		"data":     createdBook,
	})
}

func ReadBook(c *gin.Context) {
	bookToRead := models.Book{
		Id:     c.Params.ByName("id"),
		Name:   "",
		Author: "",
		Resume: "",
	}

	readedBook, err := es.ReadBook(bookToRead)

	// handle unvalid ID
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"_message": fmt.Sprintf("Book [%s] doesn't exist. Please use a valid ID.", readedBook.Id),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"_message": fmt.Sprintf("Book [%s]", readedBook.Id),
			"data":     readedBook,
		})
	}
}

func UpdateBook(c *gin.Context) {
	bookToUpdate := models.Book{
		Id:     c.Params.ByName("id"),
		Name:   c.Query("name"),
		Author: c.Query("author"),
		Resume: c.Query("resume"),
	}

	updatedBook, err := es.UpdateBook(bookToUpdate)

	// handle unvalid ID
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
	err := es.DeleteBook(c.Params.ByName("id"))

	// handle unvalid ID
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"_message": fmt.Sprintf("Book [%s] doesn't exist. Please use a valid ID.", c.Params.ByName("id")),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"_message": fmt.Sprintf("Book [%s] deleted successfully.", c.Params.ByName("id")),
		})
	}
}
