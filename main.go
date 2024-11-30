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
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

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
	{ID: "2", Title: "Froot", Artist: "Marina and the Diamonds", Year: 2015},
	{ID: "3", Title: "I've IVE", Artist: "IVE", Year: 2023},
	{ID: "4", Title: "Silence Between Songs", Artist: "Madison Beer", Year: 2023},
	{ID: "5", Title: "2", Artist: "(G)I-DLE", Year: 2024},
}

// API REST
// GET Albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// POST
func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}
