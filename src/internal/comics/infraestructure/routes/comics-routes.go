package routes

import (
	"demo/src/internal/comics/application"
	"demo/src/internal/comics/infraestructure/database"
	"demo/src/internal/comics/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)


func RegisterComicRoutes(router *gin.Engine) {
	dbComics := database.NewMySQLComics()
	// Crear instancias de los controladores
	createComicController := controllers.NewCreateComicController(*application.NewCreateComic(dbComics))
	listComicsController := controllers.NewListComicsController(*application.NewListComics(dbComics))
	byIdComicController := controllers.NewByIdComicController(application.NewComicById(dbComics))
	updateComicController := controllers.NewUpdateComicController(*application.NewUpdateComic(dbComics))
	deleteComicController := controllers.NewDeleteComicController(*application.NewDeleteComic(dbComics))

	// Registrar las rutas
	router.POST("/comic", createComicController.Execute)
	router.GET("/comic", listComicsController.Execute)
	router.GET("/comic/:id", byIdComicController.Execute)
	router.PUT("/comic/:id", updateComicController.Execute)
	router.DELETE("/comic/:id", deleteComicController.Execute)
}