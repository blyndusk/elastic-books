package router

import (
	"github.com/blyndusk/elastic-books/api/controllers"
	"github.com/gin-gonic/gin"
)

func Setup() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "elastic-book API",
		})
	})
	esRoutes(r)
	r.Run(":3333")

}

func esRoutes(r *gin.Engine) {

	r.POST("/books/search", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"goal":   "search book by name/author/resume",
			"status": "not implemented",
		})
	})

	r.GET("/books/", controllers.CreateBook)

	r.PUT("/books/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"goal":   "[update] a book by ID",
			"status": "not implemented",
		})
	})

	r.DELETE("/books/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"goal":   "[delete] a book by ID",
			"status": "not implemented",
		})
	})

}
