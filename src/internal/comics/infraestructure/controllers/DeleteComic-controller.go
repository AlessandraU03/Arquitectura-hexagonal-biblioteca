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

func (c *DeleteComicController) Execute(g *gin.Context) {
	idStr := g.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	c.useCase.Execute(int32(id))

	// Responder con un mensaje de éxito
	g.JSON(http.StatusOK, gin.H{
		"message": "Comic eliminado con exito",
	})
}