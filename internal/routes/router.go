package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRoutes(server *gin.Engine) {
	server.GET("/hello-world", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello, Kitty!"})
	})
}
