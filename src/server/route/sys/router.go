package sys

import (
	"log"

	"github.com/gin-gonic/gin"
)

// ////////////////////////////////////////////////////////////////////////////////////////////////////////
// 视图映射
// ////////////////////////////////////////////////////////////////////////////////////////////////////////

// InitRoute 初始化路由
func InitRoute(engine *gin.Engine) {
	if engine == nil {
		log.Panicln(`engine is nil!`)
	}
	// engine.Use(jwt.JWT())
	// 用户路由
	sysApi := engine.Group("/sys")
	{
		sysApi.POST("/UserLogout", UserLogout)
		sysApi.POST("/UserAdd", UserAdd)
		sysApi.GET("/UserList", UserList)
		sysApi.POST("/UserUpdate", UserUpdate)
		sysApi.POST("/UserDelete", UserDelete)
		sysApi.GET("/AccountInfo", AccountInfo)
		sysApi.GET("/UserMsg", UserMsg)

		sysApi.GET("/GroupList", GroupList)
		sysApi.POST("/GroupAdd", GroupAdd)
		sysApi.POST("/GroupDelete", GroupDelete)
		sysApi.POST("/GroupUpdate", GroupUpdate)
		sysApi.GET("/GroupMsg", GroupMsg)
		sysApi.POST("/GroupPermUpdate", GroupPermUpdate)
		sysApi.GET("/GroupPermission", GroupPermission)

		sysApi.GET("/PermissionList", PermissionList)
		sysApi.GET("/PermissionMsg", PermissionMsg)
		sysApi.POST("/PermissionAdd", PermissionAdd)
		sysApi.POST("/PermissionDelete", PermissionDelete)
		sysApi.POST("/PermissionUpdate", PermissionUpdate)

	}
}
