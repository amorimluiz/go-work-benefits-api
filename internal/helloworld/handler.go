package helloworld

import "github.com/gin-gonic/gin"

type handler struct {
}

func newHandler() *handler {
	return &handler{}
}

var HelloWorldHandler = newHandler()

func (h *handler) HelloWorld(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Hello World!"})
}
