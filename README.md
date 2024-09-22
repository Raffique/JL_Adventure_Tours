Here’s a step-by-step guide to creating a Go API using the Gin framework, GORM for ORM, and PostgreSQL from a Docker container for CRUD operations. I'll walk through setting up the environment, writing the code, and running everything in Docker.
Prerequisites

    Install Docker
    Install Go on your system (if not installed)

1. Set up PostgreSQL in Docker

Start by pulling a PostgreSQL Docker image and running a container.

```bash
docker pull postgres
docker run --name go-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=user -e POSTGRES_DB=JLAdventuresDB -p 5432:5432 -d postgres
```

This will create a PostgreSQL instance running on localhost:5432, with the following credentials:

    Database name: exampledb
    Username: user
    Password: password

2. Create the Go Project

Start by creating a new directory for your Go project and initialize a Go module.

```bash
mkdir go-gin-gorm-api
cd go-gin-gorm-api
go mod init github.com/yourusername/go-gin-gorm-api
```

3. Install the Required Go Packages

Install the necessary packages for Gin, GORM, and PostgreSQL.

```bash
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

4. Create a main.go file
```code
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
	dsn := "host=localhost user=user password=password dbname=exampledb port=5432 sslmode=disable"
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

```

5. Run PostgreSQL in Docker

Make sure your PostgreSQL container is running.

```bash
docker ps
```

If it is not running, start it with:
```bash
docker start go-postgres
```

6. Running the Go API
Run the Go application by executing the following command:
```bash
go run main.go
```
This will start the API server on localhost:8080.

7. Test the API with curl or Postman
Create a Book:
```bash
curl -X POST http://localhost:8080/books -H "Content-Type: application/json" -d '{"title": "The Hobbit", "author": "J.R.R. Tolkien"}'
```

Get All Books:
```bash
curl http://localhost:8080/books
```

Get a Book by ID:
```bash
curl http://localhost:8080/books/1
```

Update a Book:
```bash
curl -X PUT http://localhost:8080/books/1 -H "Content-Type: application/json" -d '{"title": "The Lord of the Rings", "author": "J.R.R. Tolkien"}'
```

Delete a Book:
```bash
curl -X DELETE http://localhost:8080/books/1
```

8. Dockerize the Go API

Create a Dockerfile to build the Go API as a Docker container.

Dockerfile

# Use the official Golang image
FROM golang:1.18

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

9. Build and Run the Docker Container

    Build the Docker image:
```bash
docker build -t go-gin-gorm-api .
```

    Run the Docker container:
```bash
docker run -d -p 8080:8080 --name go-gin-api --link go-postgres:postgres go-gin-gorm-api
```

The API should now be accessible at localhost:8080 and connected to the PostgreSQL database in Docker.
10. Clean Up

Stop and remove the Docker containers when done:
```bash
docker stop go-gin-api go-postgres
docker rm go-gin-api go-postgres
```

That’s it! You've set up a Go API using Gin and GORM, connected it to a PostgreSQL database running in a Docker container, and even Dockerized the API itself!