package routers

import (
	"Blog/global"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	apiRouterGroup := router.Group("api")
	// 系统配置api
	routerGroupApp := RouterGroup{apiRouterGroup}
	routerGroupApp.SettingsRouter()

	return router
}
