package routes

import (
	"demo/src/internal/books/application"
	"demo/src/internal/books/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(router *gin.Engine, createBookUseCase *application.CreateBook, listBooksUseCase *application.ListBooks, updateBookUseCase *application.UpdateBook, deleteBookUseCase *application.DeleteBook) {
	// Crear instancias de los controladores
	createBookController := controllers.NewCreateBookController(*createBookUseCase)
	listBooksController := controllers.NewListBooksController(*listBooksUseCase)
	updateBookController := controllers.NewUpdateBookController(*updateBookUseCase)
	deleteBookController := controllers.NewDeleteBookController(*deleteBookUseCase)

	// Registrar las rutas
	router.POST("/books", createBookController.Handle)
	router.GET("/books", listBooksController.Handle)
	router.PUT("/books/:id", updateBookController.Handle)
	router.DELETE("/books/:id", deleteBookController.Handle)
}