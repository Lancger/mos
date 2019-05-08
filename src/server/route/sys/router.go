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
	apiUser := engine.Group("/sys")
	{
		apiUser.POST("/UserLogout", UserLogout)
		apiUser.POST("/UserAdd", UserAdd)
		apiUser.GET("/UserList", UserList)
		apiUser.POST("/UserUpdate", UserUpdate)
		apiUser.POST("/UserDelete", UserDelete)
		apiUser.GET("/AccountInfo", AccountInfo)
		apiUser.GET("/UserMsg", UserMsg)

		apiUser.GET("/GroupList", GroupList)
		apiUser.POST("/GroupAdd", GroupAdd)
		apiUser.POST("/GroupDelete", GroupDelete)
		apiUser.POST("/GroupUpdate", GroupUpdate)

		apiUser.GET("/GroupMsg", GroupMsg)

	}
}
