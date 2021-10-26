package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album {
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	// initialize a gin router using default 
	router := gin.Default()
	// we are passing the name of the getAlbum function. This is different from passing the result of the function which would be getAlbums()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	// use the run function to attach the router to an http.Server and start the server
	router.Run()
}


// creates JSON from the slice of album structs, writing the JSON into the response
func getAlbums(c *gin.Context) {
	// call Context.IndentedJSON to serialize the struct into JSON and add it to the response
	// the first argument is the HTTP status code you want to send to the client
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the request body to newAlbum
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	// use Context.Param to retrieve the id path parameter from the URL. When you map this handler to a path, you'll include a placeholder for the parameter in the path 
    id := c.Param("id")

    // Loop over the list of albums, looking for an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more. 
// This is different from Go's built in context package
