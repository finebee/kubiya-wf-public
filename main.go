package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() http.Handler {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "hello " + name,
		})
	})

	return r
}
