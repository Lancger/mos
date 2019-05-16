package ticket

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
	ticketAPI := engine.Group("/ticket")
	{
		ticketAPI.POST("/TicketTypeAdd", TicketTypeAdd)
		ticketAPI.POST("/TicketTypeUpdate", TicketTypeUpdate)
		ticketAPI.POST("/TicketTypeDelete", TicketTypeDelete)
		ticketAPI.GET("/TicketTypeList", TicketTypeList)
		ticketAPI.GET("/TicketTypeMsg", TicketTypeMsg)

		ticketAPI.POST("/TicketSourceAdd", TicketSourceAdd)
		ticketAPI.POST("/TicketSourceUpdate", TicketSourceUpdate)
		ticketAPI.POST("/TicketSourceDelete", TicketSourceDelete)
		ticketAPI.GET("/TicketSourceList", TicketSourceList)
		ticketAPI.GET("/TicketSourceMsg", TicketSourceMsg)

		ticketAPI.POST("/TicketLevelAdd", TicketLevelAdd)
		ticketAPI.POST("/TicketLevelUpdate", TicketLevelUpdate)
		ticketAPI.POST("/TicketLevelDelete", TicketLevelDelete)
		ticketAPI.GET("/TicketLevelList", TicketLevelList)

		ticketAPI.POST("/TicketList", TicketList)
		ticketAPI.POST("/TicketAdd", TicketAdd)
		ticketAPI.POST("/TicketUpdate", TicketUpdate)
		ticketAPI.POST("/TicketCtl", TicketCtl)

		ticketAPI.POST("/TicketDelete", TicketDelete)
		ticketAPI.GET("/TicketMsg", TicketMsg)

		ticketAPI.POST("/TicketSend", TicketSend)

	}
}
