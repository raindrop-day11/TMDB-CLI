package routes

import (
	"tmdb-cli-tool/app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(router *gin.Engine) {
	tmdb := router.Group("/TMDB")
	{
		movie := new(controllers.MovieControllers)
		tmdb.GET("/movie", movie.GetMovies)
	}
}
