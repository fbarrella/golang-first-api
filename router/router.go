package router

import (
    "github.com/gin-gonic/gin"
    "first_api/controller"
)

func StartRouter() {
    router := gin.Default()
    router.GET("/albums", controller.GetAlbums)
    router.Run("localhost:7074")
}
