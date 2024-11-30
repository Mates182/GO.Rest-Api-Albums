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

// Album structure
type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

// Album data
var albums = []album{
	{ID: "1", Title: "Born to Die", Artist: "Lana del Rey", Year: 2012},
	{ID: "2", Title: "Born to Die", Artist: "Lana del Rey", Year: 2012},
	{ID: "3", Title: "Born to Die", Artist: "Lana del Rey", Year: 2012},
}
