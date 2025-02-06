package controllers

import (
	"demo/src/internal/books/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ByIdBookController struct {
	useCase *application.BookById
}

func NewByIdBookController(useCase *application.BookById) *ByIdBookController {
	return &ByIdBookController{useCase: useCase}
}

func (c *ByIdBookController) Execute(g *gin.Context) {
    idStr := g.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    // Llamamos al caso de uso para obtener el libro
    book, err := c.useCase.Execute(int32(id))
    if err != nil {
        g.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    // Respondemos con los datos del libro
    g.JSON(http.StatusOK, book)
}
