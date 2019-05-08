package iaas

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func test(ctx *gin.Context) {
	s := map[string]interface{}{}
	bodyString := `{
		"id": 1,
		"name": "jack",
		"class": {
			"name": "class6"
			"time": "20190310"
			"details": {
				"teacher": "lucy"
			}
		}
	}`
	// var resp Response
	err := json.Unmarshal([]byte(bodyString), &s)
	if err != nil {
		panic(err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "oj",
		"data":    s,
	})
}

func InitRoute(engine *gin.Engine) {
	if engine == nil {
		log.Panicln(`engine is nil!`)
	}
	route := engine.Group("/iaas/")
	route.GET("/test", test)
}

type Student struct {
	Name string `json:"name"`
}
