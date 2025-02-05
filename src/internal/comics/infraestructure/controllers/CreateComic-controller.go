package controllers

import (
	"demo/src/internal/comics/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateComicController struct {
	useCase application.CreateComic
}

func NewCreateComicController(useCase application.CreateComic) *CreateComicController {
	return &CreateComicController{useCase: useCase}
}

func (c *CreateComicController) Execute(g *gin.Context) {
	var newComic struct {
		Name      string `json:"name"`
		Autor     string `json:"autor"`
		Editorial string `json:"editorial"`
	}

	if err := g.ShouldBindJSON(&newComic); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	c.useCase.Execute(newComic.Name, newComic.Autor, newComic.Editorial)
	// Responder con el mensaje y los datos del libro creado
	g.JSON(http.StatusCreated, gin.H{
		"message":   "Comic creado con exito",
		"name":      newComic.Name,
		"autor":     newComic.Autor,
		"editorial": newComic.Editorial,
	})
}