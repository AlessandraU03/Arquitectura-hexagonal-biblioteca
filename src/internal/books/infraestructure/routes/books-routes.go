package routes

import (
	"demo/src/internal/books/application"
	"demo/src/internal/books/infraestructure/database"
	"demo/src/internal/books/infraestructure/controllers"
	"github.com/gin-gonic/gin"

)

func RegisterBookRoutes(router *gin.Engine) {
    dbBooks := database.NewMySQLBooks()
    // Crear instancias de los controladores
    createBookController := controllers.NewCreateBookController(*application.NewCreateBook(dbBooks))
    byIdBookController := controllers.NewByIdBookController(application.NewBookById(dbBooks))
    listBooksController := controllers.NewListBooksController(*application.NewListBook(dbBooks))
    updateBookController := controllers.NewUpdateBookController(*application.NewUpdateBook(dbBooks))
    deleteBookController := controllers.NewDeleteBookController(*application.NewDeleteBook(dbBooks))

    // Registrar las rutas
    router.POST("/book", createBookController.Execute)
    router.GET("/book", listBooksController.Execute)
    router.GET("/book/:id", byIdBookController.Execute)  // Ruta para obtener el libro por ID
    router.PUT("/book/:id", updateBookController.Execute)
    router.DELETE("/book/:id", deleteBookController.Execute)
}
