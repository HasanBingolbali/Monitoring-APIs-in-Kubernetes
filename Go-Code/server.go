package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Book represents the model for a book
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func init() {
	books = append(books, Book{ID: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"})
	books = append(books, Book{ID: 2, Title: "1984", Author: "George Orwell"})
}

func main() {
	r := gin.Default()

	r.GET("/books", getBooks)
	r.POST("/books", createBook)
	r.DELETE("/books/:id", deleteBook)
	r.PUT("/books/:id", updateBook)

	r.Run(":8080")
}

func getBooks(c *gin.Context) {
	sleep(100)
	c.JSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	sleep(100)
	c.JSON(http.StatusCreated, gin.H{"message": "Book added!"})
}

func deleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			sleep(100)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted!"})
			return
		}
	}
	sleepError(10)
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found."})
}

func updateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedBook Book
	if err := c.BindJSON(&updatedBook); err != nil {
		return
	}
	for i, b := range books {
		if b.ID == id {
			books[i] = updatedBook
			sleep(100)
			c.JSON(http.StatusOK, gin.H{"message": "Book updated!"})
			return
		}
	}
	sleepError(10)
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found."})
}

func sleep(ms int) {
	rand.Seed(time.Now().UnixNano())
	now := time.Now()
	n := rand.Intn(ms + now.Second()*2)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

func sleepError(ms int) {
	rand.Seed(time.Now().UnixNano())
	now := time.Now()
	n := rand.Intn(ms + now.Second())
	time.Sleep(time.Duration(n) * time.Millisecond)
}
