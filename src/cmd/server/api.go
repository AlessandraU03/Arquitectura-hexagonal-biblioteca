package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"demo/src/internal/books/application"
	"demo/src/internal/books/infraestructure/database"
	"demo/src/internal/books/infraestructure/routes"
)

func Api() {
	// Inicializar la base de datos
	db := database.NewMySQL()

	// Inicializar los casos de uso de libros
	createBookUseCase := application.NewCreateBook(db)
	listBooksUseCase := application.NewListBook(db)
	updateBookUseCase := application.NewUpdateBook(db)
	deleteBookUseCase := application.NewDeleteBook(db)

	// Configurar el router de Gin
	router := gin.Default()

	// Registrar rutas
	routes.RegisterBookRoutes(router, createBookUseCase, listBooksUseCase, updateBookUseCase, deleteBookUseCase)

	// Iniciar el servidor en el puerto 8080
	log.Println("Servidor escuchando en el puerto 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error iniciando el servidor: %v", err)
	}
}
