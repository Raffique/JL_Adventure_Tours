package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Define the model
type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var db *gorm.DB
var err error

func main() {
	// Set up database connection string
	dsn := "host=localhost user=user password=password dbname=JLAdventuresDB port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	// Auto migrate the Book model
	db.AutoMigrate(&Book{})

	// Initialize Gin router
	router := gin.Default()

	// Define routes
	router.POST("/books", createBook)
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBook)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)

	// Start server
	router.Run(":8080")
}

// Handlers for CRUD operations

// Create a new book
func createBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&book)
	c.JSON(http.StatusOK, book)
}

// Get all books
func getBooks(c *gin.Context) {
	var books []Book
	db.Find(&books)
	c.JSON(http.StatusOK, books)
}

// Get a specific book by ID
func getBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// Update a book
func updateBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&book)
	c.JSON(http.StatusOK, book)
}

// Delete a book
func deleteBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
