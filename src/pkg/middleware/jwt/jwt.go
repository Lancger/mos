package jwt

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"mos/src/pkg/e"
	"mos/src/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		code = e.SUCCESS
		token := ctx.Request.Header.Get("X-Token")
		log.Printf(token)
		if token == `` {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    e.TOKEN_EMPTY,
				"message": e.GetCodeMessage(e.TOKEN_EMPTY),
				"data":    "",
			})
			ctx.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			code = e.TOKEN_INVAILD
			log.Println(err, claims)
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.TOKEN_EXPRIED
		}
		if code != e.SUCCESS {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": e.GetCodeMessage(code),
				"data":    "",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
