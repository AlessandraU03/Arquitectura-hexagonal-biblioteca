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

func (c *ListComicsController) Handle(ctx *gin.Context) {
	ComiListComics, err := c.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve ComiListComics"})
		return
	}

	ctx.JSON(http.StatusOK, ComiListComics)
}