package controllers

import (
	"demo/src/internal/books/application"
	"demo/src/internal/books/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateBookController struct {
	useCase application.CreateBook
}

func NewCreateBookController(useCase application.CreateBook) *CreateBookController {
	return &CreateBookController{useCase: useCase}
}

func (c *CreateBookController) Handle(ctx *gin.Context) {
	var newBook struct {
		Name      string `json:"name"`
		Autor     string `json:"autor"`
		Categoria string `json:"categoria"`
	}

	// Validar los datos enviados en el cuerpo de la solicitud
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Crear el libro usando el caso de uso
	book := entities.NewBook(newBook.Name, newBook.Autor, newBook.Categoria)
	if err := c.useCase.Execute(book); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	// Responder con el mensaje y los datos del libro creado
	ctx.JSON(http.StatusCreated, gin.H{
		"name":      book.Name,
		"autor":     book.Autor,
		"categoria": book.Categoria,
	})
}