package controllers

import (
	"demo/src/internal/comics/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ByIdComicController struct {
	useCase *application.ComicById
}

func NewByIdComicController(useCase *application.ComicById) *ByIdComicController {
	return &ByIdComicController{useCase: useCase}
}

func (c *ByIdComicController) Execute(g *gin.Context) {
    idStr := g.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    // Llamamos al caso de uso para obtener el libro
    Comic, err := c.useCase.Execute(int32(id))
    if err != nil {
        g.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    // Respondemos con los datos del libro
    g.JSON(http.StatusOK, Comic)
}
