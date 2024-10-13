package controllers

import (
	"net/http"
	"strconv"

	"github.com/Gabs-Leo/FATEC-Go-CRUD/models"
	"github.com/gin-gonic/gin"
)

func GetBooks(context *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	context.JSON(http.StatusOK, books)
}

func GetBook(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var book models.Book
	if err := models.DB.First(&book, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	context.JSON(http.StatusOK, book)
}

func CreateBook(context *gin.Context) {
	var book models.Book
	if err := context.ShouldBindJSON(&book); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&book)
	context.JSON(http.StatusCreated, book)
}

func UpdateBook(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var book models.Book
	if err := models.DB.First(&book, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := context.ShouldBindJSON(&book); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Save(&book)
	context.JSON(http.StatusOK, book)
}

func DeleteBook(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var book models.Book
	if err := models.DB.First(&book, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	models.DB.Delete(&book)
	context.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
