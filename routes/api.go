package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(router *gin.Engine) {
	router.Group("/TMDB")
	{
		router.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "world",
			})
		})
	}
}
