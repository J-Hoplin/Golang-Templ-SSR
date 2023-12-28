package internal

import (
	"cicd/internal/database"
	"cicd/internal/todo"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func InitServer() {
	root := gin.Default()

	// Load Dot env
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Fail to load .env file:", err.Error())
	}

	// Initialize Database
	database.Connect()
	defer database.Disconnect()

	// TODO Handler
	todoRoute := root.Group("/todo")
	todo.EnrollRoute(todoRoute)

	// When something makes an error while bootstraping server
	if err := root.Run(":8080"); err != nil {
		log.Fatalln("Fail to start server:", err.Error())
	}

}
