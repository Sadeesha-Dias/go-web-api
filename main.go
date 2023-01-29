package main

import (
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

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func addBooks(p *gin.Context) {
	var newBook book
	if err := p.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	p.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()

	//GET
	router.GET("/books", getBooks)

	//POST
	router.POST("/addbooks", addBooks)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
