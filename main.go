package main

import (
	"tmdb-cli-tool/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	//初始化路由
	bootstrap.SetupRpute(router)

	router.Run(":8000")
}
