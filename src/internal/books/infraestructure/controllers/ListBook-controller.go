package controllers

import (
	"demo/src/internal/books/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListBooksController struct {
	useCase application.ListBooks
}

func NewListBooksController(useCase application.ListBooks) *ListBooksController {
	return &ListBooksController{useCase: useCase}
}

func (c *ListBooksController) Execute(g *gin.Context) {
	books, err := c.useCase.Execute()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}

	g.JSON(http.StatusOK, books)
}