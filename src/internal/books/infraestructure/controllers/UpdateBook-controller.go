package controllers

import (
	"demo/src/internal/books/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateBookController struct {
	useCase application.UpdateBook
}

func NewUpdateBookController(useCase application.UpdateBook) *UpdateBookController {
	return &UpdateBookController{useCase: useCase}
}

func (c *UpdateBookController) Execute(g *gin.Context) {
	var updatedBook struct {
		Name      string `json:"name"`
		Autor     string `json:"autor"`
		Categoria string `json:"categoria"`
	}

	// Validar los datos enviados en el cuerpo de la solicitud
	if err := g.ShouldBindJSON(&updatedBook); err != nil {
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

	// Ejecutar el caso de uso para actualizar el libro
	c.useCase.Execute(int32(id), updatedBook.Name, updatedBook.Autor, updatedBook.Categoria)

	// Responder con los datos del libro actualizado
	g.JSON(http.StatusOK, gin.H{
			"message": "Libro actualizado con exito",
			"id":        id,
			"name":      updatedBook.Name,
			"autor":     updatedBook.Autor,
			"categoria": updatedBook.Categoria,
	})
}
