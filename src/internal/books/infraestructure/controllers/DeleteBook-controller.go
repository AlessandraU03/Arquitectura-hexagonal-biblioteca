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

func (c *DeleteBookController) Handle(ctx *gin.Context) {
	// Obtener el ID del libro desde la URL
	id := ctx.Param("id")
	idInt32, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Ejecutar el caso de uso para eliminar el libro
	if err := c.useCase.Execute(int32(idInt32)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	// Responder con un mensaje de éxito
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book deleted",
	})
}