package controllers

import (
	"net/http"

	"github.com/blyndusk/elastic-books/api/es"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": es.Esclient.String(),
	})
}