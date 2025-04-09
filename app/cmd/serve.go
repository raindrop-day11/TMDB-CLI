package cmd

import (
	"tmdb-cli-tool/bootstrap"
	"tmdb-cli-tool/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var Serve = &cobra.Command{
	Use: "serve",
	Run: runWeb,
}

func runWeb(cmd *cobra.Command, args []string) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	//初始化路由
	bootstrap.SetupRpute(router)

	router.Run(":" + config.GetString("app.port"))
}
