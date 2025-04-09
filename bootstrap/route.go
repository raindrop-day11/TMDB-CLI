package bootstrap

import (
	"net/http"
	"strings"
	"tmdb-cli-tool/routes"

	"github.com/gin-gonic/gin"
)

func SetupRpute(router *gin.Engine) {
	//配置中间件
	RegisterGlobalsMiddlewares(router)
	//注册路由信息
	routes.RegisterAPIRoutes(router)
	//处理未命名路由

}

func RegisterGlobalsMiddlewares(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func Handle404Route(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		//获取请求头的accept信息
		acceptString := c.Request.Header.Get("Accept")

		//如果是HTML的话
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面不存在")
		} else {
			//默认返回的是JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "页面不存在",
			})
		}
	})
}
