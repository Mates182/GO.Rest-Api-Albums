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
	// Index page
	router.GET("/", handler)

	// RESTful routes
	router.GET("/albums", getAlbums)            //Get all albums
	router.GET("albums/:id", getAlbumByID)      //Get album by ID
	router.POST("/albums", postAlbums)          //Add a new album
	router.PUT("/albums/:id", updateAlbumByID)  //Update an album by ID
	router.DELETE("/albums/:id", deleteAlbum)   //Delete an album by ID
	router.PATCH("/albums/:id", patchAlbumByID) //Partially update an album by ID

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
// GET: Get all Albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// GET: Get album by ID
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

// POST: Add a new album
func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return

	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}

// PUT: Update album by ID
func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, a := range albums {
		if a.ID == id {
			albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusOK, updatedAlbum)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

// DELETE: Delete an album by id
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")
	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.JSON(http.StatusOK, albums)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

// PATCH: Partially update an album by ID
func patchAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var updates map[string]interface{}
	if err := c.BindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, a := range albums {
		if a.ID == id {
			if title, ok := updates["title"].(string); ok {
				albums[i].Title = title
			}
			if artist, ok := updates["artist"].(string); ok {
				albums[i].Artist = artist
			}
			if year, ok := updates["year"].(float64); ok {
				albums[i].Year = int(year)
			}
			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}
