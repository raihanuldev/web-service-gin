package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

// fucntion for GetAllItems
func GetallItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// POST Album
func AddAlbum(c *gin.Context) {
	var newAlbm album
	err := c.BindJSON(&newAlbm)
	if err != nil {
		return
	}
	albums = append(albums, newAlbm)
	c.IndentedJSON(http.StatusCreated, albums)
}

func main() {
	mux := gin.Default()
	mux.GET("/albums", GetallItems)
	mux.POST("/albums", AddAlbum)
	mux.Run("localhost:8080")
}
