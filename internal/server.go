package internal

import (
	"cicd/internal/database"
	"cicd/internal/renderer"
	"cicd/internal/todo"
	"cicd/internal/utility"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitServer() {
	root := gin.Default()

	// Integrate Gin with Templ
	root.HTMLRender = &renderer.TemplRender{}
	// Static Assets
	root.Static("/static", "static")

	// Load Dot env
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Fail to load .env file:", err.Error())
	}
	// Initialize Database
	database.Connect()
	defer database.Disconnect()

	// Generate dummy data if it's docker environment
	if os.Getenv("ENV") == "docker" {
		GenerateDummy()
	}

	// TODO Handler
	todoRoute := root.Group("/todo")
	todo.EnrollRoute(todoRoute)

	// Utility Handler
	utilityRoute := root.Group("/utility")
	utility.EnrollRoute(utilityRoute)

	// When something makes an error while bootstraping server
	if err := root.Run(":8080"); err != nil {
		log.Fatalln("Fail to start server:", err.Error())
	}

}
