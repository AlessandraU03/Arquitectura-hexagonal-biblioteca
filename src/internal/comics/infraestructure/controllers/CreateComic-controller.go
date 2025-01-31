package controllers

import (
	"demo/src/internal/comics/application"
	"demo/src/internal/comics/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateComicController struct {
	useCase application.CreateComic
}

func NewCreateComicController(useCase application.CreateComic) *CreateComicController {
	return &CreateComicController{useCase: useCase}
}

func (c *CreateComicController) Handle(ctx *gin.Context) {
	var newComic struct {
		Name      string `json:"name"`
		Autor     string `json:"autor"`
		Editorial string `json:"editorial"`
	}

	// Validar los datos enviados en el cuerpo de la solicitud
	if err := ctx.ShouldBindJSON(&newComic); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Crear el libro usando el caso de uso
	book := entities.NewComic(newComic.Name, newComic.Autor, newComic.Editorial)
	if err := c.useCase.Execute(book); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	// Responder con el mensaje y los datos del libro creado
	ctx.JSON(http.StatusCreated, gin.H{
		"name":      book.Name,
		"autor":     book.Autor,
		"editorial": book.Editorial,
	})
}