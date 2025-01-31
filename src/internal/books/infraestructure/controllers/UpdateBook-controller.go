package controllers

import (
	"demo/src/internal/books/application"
	"demo/src/internal/books/domain/entities"
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

func (c *UpdateBookController) Handle(ctx *gin.Context) {
	var updatedBook struct {
		Name      string `json:"name"`
		Autor     string `json:"autor"`
		Categoria string `json:"categoria"`
	}

	// Validar los datos enviados en el cuerpo de la solicitud
	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
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
	book := entities.NewBook(updatedBook.Name, updatedBook.Autor, updatedBook.Categoria)
	book.ID = int32(idInt32)

	// Ejecutar el caso de uso para actualizar el libro
	if err := c.useCase.Execute(book); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	// Responder con los datos del libro actualizado
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book updated",
		"book": gin.H{
			"id":        book.ID,
			"name":      book.Name,
			"autor":     book.Autor,
			"categoria": book.Categoria,
		},
	})
}