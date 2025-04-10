package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tmdb-cli-tool/app/models/movie"
	"tmdb-cli-tool/pkg/console"

	"github.com/gin-gonic/gin"
)

type MovieControllers struct{}

func (*MovieControllers) GetMovies(c *gin.Context) {

	movieType := c.Param("movie")

	var language, page string
	if str, ok := c.GetQuery("language"); ok {
		language = str
	}

	if str, ok := c.GetQuery("page"); ok {
		page = str
	}

	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?language=%s&page=%s", movieType, language, page)

	data := movie.GetMovies(url)

	if movieType == "now_playing" || movieType == "upcoming" {
		var movies movie.MoviesWithDate
		err := json.Unmarshal(data, &movies)
		console.ExitIf(err)

		c.JSON(http.StatusOK, gin.H{
			"messages": movies,
		})
	} else if movieType == "top_rated" || movieType == "popular" {
		var movies movie.MoviesWithoutDate
		err := json.Unmarshal(data, &movies)
		console.ExitIf(err)

		c.JSON(http.StatusOK, gin.H{
			"messages": movies,
		})
	} else {
		c.String(http.StatusBadRequest, "电影类型错误，只有now_playing,upcoming,top_rated,popular四种")
	}
}
