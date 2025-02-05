package controllers

import (
	"demo/src/internal/books/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteBookController struct {
	useCase application.DeleteBook
}

func NewDeleteBookController(useCase application.DeleteBook) *DeleteBookController {
	return &DeleteBookController{useCase: useCase}
}

func (c *DeleteBookController) Execute(g *gin.Context) {
	// Obtener el ID del libro desde la URL
	idStr := g.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Ejecutar el caso de uso para eliminar el libro
	c.useCase.Execute(int32(id))

	// Responder con un mensaje de éxito
	g.JSON(http.StatusOK, gin.H{
		"message": "Libro eliminado con exito",
	})
}
