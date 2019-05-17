package jenkins

import (
	"fmt"
	"mos/src/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JenkinsPost(ctx *gin.Context) {
	var (
		req JobPost
		j   JenkinsJob
	)
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": "参数错误," + fmt.Sprintf("错误信息: %s", err),
		})
		return
	}
	// fmt.Printf("%+vn", req)
	if status, err := j.Create(&req); status == false {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": "上报失败," + fmt.Sprintf("错误信息: %s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "上报成功",
	})
	return
}
