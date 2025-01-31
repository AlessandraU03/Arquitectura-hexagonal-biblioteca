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

func (c *ListBooksController) Handle(ctx *gin.Context) {
	books, err := c.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}

	ctx.JSON(http.StatusOK, books)
}