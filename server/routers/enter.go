package routers

import (
	"github.com/gin-gonic/gin"
)

// RouterGroup 结构体扩展了gin的RouterGroup，用于方便地组织路由。
type RouterGroup struct {
	*gin.RouterGroup
}

// InitRouter 初始化路由，返回gin的引擎实例。
// 它负责设置gin的运行模式，并初始化默认的路由。
// 通过路由分组和分层，实现了对API的组织。
// 返回值: *gin.Engine - gin的引擎实例，用于启动HTTP服务。
func InitRouter() *gin.Engine {
	gin.SetMode("release")
	router := gin.Default() // 创建默认的gin路由器。

	// 路由分组 - 将所有api路由分组到/api路径下。
	apiRouterGroup := router.Group("api")

	// 路由分层 - 通过RouterGroup类型组织路由，使得路由的设置更加模块化和层次化。
	// 系统配置api - 实例化RouterGroup，并注册设置和图片的相关路由。
	routerGroupApp := RouterGroup{apiRouterGroup}
	routerGroupApp.SettingsRouter()
	routerGroupApp.ImagesRouter()

	return router // 返回配置好的gin引擎实例。
}
