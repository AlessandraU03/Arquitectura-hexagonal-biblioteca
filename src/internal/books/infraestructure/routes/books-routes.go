package routes

import (
	"demo/src/internal/books/application"
	"demo/src/internal/books/domain/entities"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// RegisterBookRoutes registra las rutas para los libros.
func RegisterBookRoutes(router *gin.Engine, createBookUseCase *application.CreateBook, listBooksUseCase *application.ListBooks, updateBookUseCase *application.UpdateBook, deleteBookUseCase *application.DeleteBook) {

	// Ruta para crear un libro
	router.POST("/books", func(c *gin.Context) {
		var newBook struct {
			Name      string `json:"name"`
			Autor     string `json:"autor"`
			Categoria string `json:"categoria"`
		}

		// Validar los datos enviados en el cuerpo de la solicitud
		if err := c.ShouldBindJSON(&newBook); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		// Crear el libro usando el caso de uso
		Book := entities.NewBook(newBook.Name, newBook.Autor, newBook.Categoria)
		err := createBookUseCase.Execute(Book)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create book"})
			return
		}

		// Responder con el mensaje y los datos del libro creado
		c.JSON(201, gin.H{
			"name":     Book.Name,
			"autor":    Book.Autor,
			"categoria": Book.Categoria,
		})
	})

	// Ruta para listar todos los libros
	router.GET("/books", func(c *gin.Context) {
		books, err := listBooksUseCase.Execute()
		if err != nil {
			log.Printf("Error retrieving books: %v", err)
			c.JSON(500, gin.H{"error": "Failed to retrieve books"})
			return
		}

		c.JSON(200, books)
	})

	// Ruta para actualizar un libro
	router.PUT("/books/:id", func(c *gin.Context) {
		var updatedBook struct {
			Name      string `json:"name"`
			Autor     string `json:"autor"`
			Categoria string `json:"categoria"`
		}

		// Validar los datos enviados en el cuerpo de la solicitud
		if err := c.ShouldBindJSON(&updatedBook); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		// Obtener el ID del libro desde la URL
		id := c.Param("id")

		// Convertir el ID a int32
		idInt32, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID format"})
			return
		}

		Book := entities.NewBook(updatedBook.Name, updatedBook.Autor, updatedBook.Categoria)
		Book.ID = int32(idInt32) // Asignar el ID al libro

		// Ejecutar el caso de uso para actualizar el libro
		err = updateBookUseCase.Update(Book)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update book"})
			return
		}

		// Responder con los datos del libro actualizado
		c.JSON(200, gin.H{
			"message": "Book updated",
			"book": gin.H{
				"id":       Book.ID,
				"name":     Book.Name,
				"autor":    Book.Autor,
				"categoria": Book.Categoria,
			},
		})
	})

	// Ruta para eliminar un libro
	router.DELETE("/books/:id", func(c *gin.Context) {
		// Obtener el ID del libro desde la URL (convertido a int32)
		id := c.Param("id")

		// Convertir el ID a int32
		idInt32, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID format"})
			return
		}

		// Ejecutar el caso de uso para eliminar el libro
		err = deleteBookUseCase.Delete(int32(idInt32))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete book"})
			return
		}

		// Responder con un mensaje de éxito
		c.JSON(200, gin.H{
			"message": "Book deleted",
		})
	})
}
