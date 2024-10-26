package routes

import (
	"github.com/amorimluiz/work_benefits_api/internal/helloworld"
	"github.com/gin-gonic/gin"
)

func SetRoutes(server *gin.Engine) {
	server.GET("/helloworld", helloworld.HelloWorldHandler.HelloWorld)
}
