package controllers

import (
	"demo/src/internal/comics/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteComicController struct {
	useCase application.DeleteComic
}

func NewDeleteComicController(useCase application.DeleteComic) *DeleteComicController {
	return &DeleteComicController{useCase: useCase}
}

func (c *DeleteComicController) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt32, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Ejecutar el caso de uso para eliminar el libro
	if err := c.useCase.Execute(int32(idInt32)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Comic"})
		return
	}

	// Responder con un mensaje de éxito
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Comic deleted",
	})
}