package main

import (
	"os"

	"github.com/amorimluiz/work_benefits_api/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	server := gin.Default()

	routes.SetRoutes(server)

	server.Run(":" + os.Getenv("APP_PORT"))
}
