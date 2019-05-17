package project

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
	projectAPI := engine.Group("/project")
	{
		projectAPI.POST("/ProjectAdd", ProjectAdd)
		projectAPI.POST("/ProjectUpdate", ProjectUpdate)
		projectAPI.POST("/ProjectDelete", ProjectDelete)

		projectAPI.GET("/ProjectList", ProjectList)
		projectAPI.GET("/ProjectMsg", ProjectMsg)

	}
}
