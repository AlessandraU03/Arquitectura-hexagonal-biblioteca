package controllers

import (
	"demo/src/internal/comics/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListComicsController struct {
	useCase application.ListComics
}

func NewListComicsController(useCase application.ListComics) *ListComicsController {
	return &ListComicsController{useCase: useCase}
}

func (c *ListComicsController) Execute(g *gin.Context) {
	comics, err := c.useCase.Execute()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve ComiListComics"})
		return
	}

	g.JSON(http.StatusOK, comics)
}