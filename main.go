package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "GO REST API",
	})
}
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", handler)
	router.Run("localhost:8080")
}
