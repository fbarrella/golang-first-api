package controller

import (
    "first_api/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

var albums = []models.Album{
    {ID: "1", Title: "Imagine", Artist: "John Lennon", Price: 55.99},
    {ID: "2", Title: "White Album", Artist: "The Beatles", Price: 59.99},
}

func GetAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}
