// https://go.dev/doc/tutorial/web-service-gin
package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
// Struct tags such as json:"artist" specify what a field’s name should be
// when the struct’s contents are serialized into JSON. Without them, the JSON
// would use the struct’s capitalized field names – a style not as common in JSON.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
// gin.Context is the most important part of Gin. It carries request details,
// validates and serializes JSON, and more.
func getAlbums(c *gin.Context) {
	log.Printf("Client IP: %v", c.ClientIP())
	// gin.Context.IndentedJSON serializes the struct into JSON and add it to the response
	// To send more compact (unformatted) JSON, use Context.JSON
	// http.StatusOK is the HTTP status code you want to send to the client
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		log.Fatalf("Invalid request body: %v")
		return
	}
	newAlbum.ID = strconv.Itoa(len(albums) + 1)
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")
}
