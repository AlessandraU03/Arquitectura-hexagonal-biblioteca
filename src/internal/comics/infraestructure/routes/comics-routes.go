package routes

import (
	"demo/src/internal/comics/application"
	"demo/src/internal/comics/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterComicRoutes(router *gin.Engine, createComicUseCase *application.CreateComic, listComicsUseCase *application.ListComics, updateComicUseCase *application.UpdateComic, deleteComicUseCase *application.DeleteComic) {
	// Crear instancias de los controladores
	createComicController := controllers.NewCreateComicController(*createComicUseCase)
	listComicsController := controllers.NewListComicsController(*listComicsUseCase)
	updateComicController := controllers.NewUpdateComicController(*updateComicUseCase)
	deleteComicController := controllers.NewDeleteComicController(*deleteComicUseCase)

	// Registrar las rutas
	router.POST("/comic", createComicController.Handle)
	router.GET("/comic", listComicsController.Handle)
	router.PUT("/comic/:id", updateComicController.Handle)
	router.DELETE("/comic/:id", deleteComicController.Handle)
}