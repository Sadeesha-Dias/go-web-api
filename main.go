package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "001", Title: "Quantum Physics and Wormhole Thermodynamics", Author: "Dr. Albeto Terami", Quantity: 10},
	{ID: "002", Title: "Modernity and Women", Author: "Julia Rodriguez", Quantity: 50},
	{ID: "003", Title: "Bio Coding and Human Body", Author: "Dr. Helsimo Juan & Dr. Cassie Witicker", Quantity: 25},
	{ID: "004", Title: "The Lost City of Sahara", Author: "Sandra Qualifa", Quantity: 38},
}

// get all the books ---------------------------------------
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// add a new book ------------------------------------------
func addBooks(p *gin.Context) {
	var newBook book
	if err := p.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	p.IndentedJSON(http.StatusCreated, newBook)
}

// get a book by ID ----------------------------------------
func fetchBookById(id string) (*book, error) {
	for i, element := range books {
		if element.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("Book not found!")
}

func bookById(a *gin.Context) {
	id := a.Param("id")
	book, err := fetchBookById(id)

	if err != nil {
		a.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	a.IndentedJSON(http.StatusOK, book)
}

// checkout a book --------------------------------------------
func checkOutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id query parameter"})
		return
	}

	book, err := fetchBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Insufficient book amount!"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)

}

// return a book ---------------------------------------------
func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id query parameter"})
		return
	}

	book, err := fetchBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

// main function --------------------------------------------*
func main() {
	router := gin.Default()

	//GET
	router.GET("/books", getBooks)
	router.GET("books/:id", bookById)

	//POST
	router.POST("/addbooks", addBooks)

	//PATCH
	router.PATCH("/checkout", checkOutBook)
	router.PATCH("/returnbook", returnBook)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
