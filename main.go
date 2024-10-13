package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Gabs-Leo/FATEC-Go-CRUD/controllers"
	"github.com/Gabs-Leo/FATEC-Go-CRUD/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	models.InitDatabase()
	router := gin.Default()

	// Routes
	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetBook)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
