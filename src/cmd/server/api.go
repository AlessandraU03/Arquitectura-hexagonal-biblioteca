package server

import (
	rBooks "demo/src/internal/books/infraestructure/routes"
	rComics "demo/src/internal/comics/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func Api() {
	// Crear el enrutador de Gin
	router := gin.Default()

	// Registrar rutas de books y comics
	rBooks.RegisterBookRoutes(router)
	rComics.RegisterComicRoutes(router)

	// Iniciar el servidor en el puerto 8080
	router.Run(":8080")
}
