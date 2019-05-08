package server

import (
	"fmt"
	sysutil "mos/src/server/route/sys"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Serve Run
func Serve() {
	engine := gin.Default()
	// 允许使用跨域请求  全局中间件
	engine.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	engine.POST("/InitTable", sysutil.InitTable)
	engine.POST("/UserLogin", sysutil.UserLogin)

	sysutil.InitRoute(engine)
	// apiV1 := engine.Group("/api/v1")
	// 用户组路由
	// apiV1.POST("/AddGroup", AddGroup)
	// apiV1.GET("/ListGroup", ListGroup)
	// sysutil.InitRoute(engine)
	// iaas.InitRoute(engine)

	portVal := fmt.Sprintf(`127.0.0.1:99`)
	engine.Run(portVal)
}
