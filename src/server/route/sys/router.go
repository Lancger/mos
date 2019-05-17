package sys

import (
	"log"
	"mos/src/pkg/middleware/jwt"

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
	engine.Use(jwt.JWT())
	// 用户路由
	sysAPI := engine.Group("/sys")
	{
		sysAPI.POST("/UserLogout", UserLogout)
		sysAPI.POST("/UserAdd", UserAdd)
		sysAPI.GET("/UserList", UserList)
		sysAPI.POST("/UserUpdate", UserUpdate)
		sysAPI.POST("/UserDelete", UserDelete)
		sysAPI.GET("/AccountInfo", AccountInfo)
		sysAPI.GET("/UserMsg", UserMsg)

		sysAPI.GET("/GroupList", GroupList)
		sysAPI.POST("/GroupAdd", GroupAdd)
		sysAPI.POST("/GroupDelete", GroupDelete)
		sysAPI.POST("/GroupUpdate", GroupUpdate)
		sysAPI.GET("/GroupMsg", GroupMsg)
		sysAPI.POST("/GroupPermUpdate", GroupPermUpdate)
		sysAPI.GET("/GroupPermission", GroupPermission)

		sysAPI.GET("/PermissionList", PermissionList)
		sysAPI.GET("/PermissionMsg", PermissionMsg)
		sysAPI.POST("/PermissionAdd", PermissionAdd)
		sysAPI.POST("/PermissionDelete", PermissionDelete)
		sysAPI.POST("/PermissionUpdate", PermissionUpdate)

		sysAPI.GET("/GroupUserCas", GroupUserCas)

	}
}
