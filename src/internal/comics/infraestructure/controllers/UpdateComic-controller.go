package controllers

import (
	"demo/src/internal/comics/application"
	"demo/src/internal/comics/domain/entities"
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

func (c *UpdateComicController) Handle(ctx *gin.Context) {
	var updatedComic struct {
		Name      string `json:"name"`
		Autor     string `json:"autor"`
		Editorial string `json:"editorial"`
	}

	// Validar los datos enviados en el cuerpo de la solicitud
	if err := ctx.ShouldBindJSON(&updatedComic); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Obtener el ID del libro desde la URL
	id := ctx.Param("id")
	idInt32, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Crear el libro con los datos actualizados
	Comic := entities.NewComic(updatedComic.Name, updatedComic.Autor, updatedComic.Editorial)
	Comic.Id = int32(idInt32)

	// Ejecutar el caso de uso para actualizar el libro
	if err := c.useCase.Execute(Comic); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Comic"})
		return
	}

	// Responder con los datos del libro actualizado
	ctx.JSON(http.StatusOK, Comic)
			
}