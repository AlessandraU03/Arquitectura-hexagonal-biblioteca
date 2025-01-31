package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	apBooks "demo/src/internal/books/application"
	databaseBooks "demo/src/internal/books/infraestructure/database"
	routesBooks "demo/src/internal/books/infraestructure/routes"

	apComics "demo/src/internal/comics/application"
	databaseComics "demo/src/internal/comics/infraestructure/database"
	routesComics "demo/src/internal/comics/infraestructure/routes"
)

func Api() {
	// Inicializar la base de datos de libros y comics
	dbBooks := databaseBooks.NewMySQLBooks()
	dbComics := databaseComics.NewMySQLComics()

	// Inicializar los casos de uso de libros
	createBookUseCase := apBooks.NewCreateBook(dbBooks)
	listBooksUseCase := apBooks.NewListBook(dbBooks)
	updateBookUseCase := apBooks.NewUpdateBook(dbBooks)
	deleteBookUseCase := apBooks.NewDeleteBook(dbBooks)

	// Inicializar los casos de uso de comics
	createComicUseCase := apComics.NewCreateComic(dbComics)
	listComicsUseCase := apComics.NewListComics(dbComics)
	updateComicUseCase := apComics.NewUpdateComic(dbComics)
	deleteComicUseCase := apComics.NewDeleteComic(dbComics)

	// Configurar el router de Gin
	router := gin.Default()

	// Registrar rutas de libros
	routesBooks.RegisterBookRoutes(router, createBookUseCase, listBooksUseCase, updateBookUseCase, deleteBookUseCase)

	// Registrar rutas de comics
	routesComics.RegisterComicRoutes(router, createComicUseCase, listComicsUseCase, updateComicUseCase, deleteComicUseCase)

	// Iniciar el servidor en el puerto 8080
	log.Println("Servidor escuchando en el puerto 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error iniciando el servidor: %v", err)
	}
}