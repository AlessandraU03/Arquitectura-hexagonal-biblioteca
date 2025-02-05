package controllers

import (
	"demo/src/internal/books/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateBookController struct {
	useCase application.CreateBook
}

func NewCreateBookController(useCase application.CreateBook) *CreateBookController {
	return &CreateBookController{useCase: useCase}
}

func (c *CreateBookController) Execute(g *gin.Context) {
	var newBook struct {
		Name      string `json:"name"`
		Autor     string `json:"autor"`
		Categoria string `json:"categoria"`
	}

	if err := g.ShouldBindJSON(&newBook); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	c.useCase.Execute(newBook.Name, newBook.Autor, newBook.Categoria)

	g.JSON(http.StatusCreated, gin.H{
		"message":   "Libro creado con exito",
		"name":      newBook.Name,
		"autor":     newBook.Autor,
		"categoria": newBook.Categoria,
	})
}
