package main

import (
	"tmdb-cli-tool/bootstrap"
	configbts "tmdb-cli-tool/config"
	"tmdb-cli-tool/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	configbts.InitLize()
}

func main() {
	router := gin.New()

	//初始化配置
	config.InitConfig()

	//初始化路由
	bootstrap.SetupRpute(router)

	router.Run(":" + config.GetString("app.port"))
}
