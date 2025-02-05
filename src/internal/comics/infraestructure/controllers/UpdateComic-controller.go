package controllers

import (
	"demo/src/internal/comics/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateComicController struct {
	useCase application.UpdateComic
}

func NewUpdateComicController(useCase application.UpdateComic) *UpdateComicController {
	return &UpdateComicController{useCase: useCase}
}

func (c *UpdateComicController) Execute(g *gin.Context) {
	var updatedComic struct {
		Name      string `json:"name"`
		Autor     string `json:"autor"`
		Editorial string `json:"editorial"`
	}

	// Validar los datos enviados en el cuerpo de la solicitud
	if err := g.ShouldBindJSON(&updatedComic); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Obtener el ID del libro desde la URL
	idStr := g.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	c.useCase.Execute(int32(id), updatedComic.Name, updatedComic.Autor, updatedComic.Editorial)

	g.JSON(http.StatusOK, gin.H{
		"message": "Comic actualizado con exito",
		"book": gin.H{
			"id":        id,
			"name":      updatedComic.Name,
			"autor":     updatedComic.Autor,
			"categoria": updatedComic.Editorial,
		},
	})
			
}