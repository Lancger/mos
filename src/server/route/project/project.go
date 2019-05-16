package project

import (
	"fmt"
	"mos/src/glo"
	"mos/src/glo/comfunc"
	"mos/src/pkg/e"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProjectAdd 添加工单类型
func ProjectAdd(ctx *gin.Context) {
	type reqPostData struct {
		Name string `json:"name"`
	}
	var (
		r reqPostData
	)
	if err := ctx.BindJSON(&r); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": e.PARAM_ERROR_MSG,
		})
		return
	}
	t := Project{
		Name: r.Name,
	}
	if err := glo.Db.Create(&t).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "添加工单类型失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "添加工单类型成功",
	})
	return
}

// ProjectList 工单类型列表
func ProjectList(ctx *gin.Context) {
	// 初始化请求参数变量
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize = comfunc.GetDefaultPageSize(pageSize)
	page = comfunc.GetDefaultPage(page)
	// init params
	querySet := make([]Project, 0)
	res := make([]ProjectInfo, 0)
	dbQuery := glo.Db.Model(&Project{})
	var (
		typeName string
		total    int
	)
	// 设置查询用户名
	typeName = ctx.Query("search")
	if typeName != `` {
		// 根据groupname模糊查询，可以将gorm链接添加条件后，赋值覆盖自身，得到不定条件的链式查询效果
		dbQuery = dbQuery.Where("name LIKE ?", fmt.Sprintf("%%%s%%", typeName))
	}

	if err := dbQuery.Limit(pageSize).Offset((page - 1) * pageSize).Order("id desc").Find(&querySet).Count(&total).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "查询失败 " + err.Error(),
		})
		return
	}
	for _, r := range querySet {
		var g ProjectInfo
		g.ID = r.ID
		g.Name = r.Name
		g.CreateTime = comfunc.FormatTs(r.CreatedAt.Unix())
		res = append(res, g)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":       e.SUCCESS,
		"message":    "success",
		"data":       res,
		"total":      total,
		"total_page": comfunc.FlorPageInt(pageSize, total),
		"page_size":  pageSize,
		"page":       page,
	})
	return
}

// ProjectUpdate 更新工单类型
func ProjectUpdate(ctx *gin.Context) {
	type reqPostData struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	var reqData reqPostData
	if err := ctx.BindJSON(&reqData); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": e.PARAM_ERROR_MSG,
		})
		return
	}
	if err := glo.Db.Model(&Project{}).Where("id = ?", reqData.ID).Update(&Project{Name: reqData.Name}).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "更新失败:" + fmt.Sprintf("%s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "更新成功",
	})
	return
}

// ProjectUpdate 更新工单类型
func ProjectDelete(ctx *gin.Context) {
	type reqPostData struct {
		ID uint `json:"id"`
	}
	var reqData reqPostData
	if err := ctx.BindJSON(&reqData); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": e.PARAM_ERROR_MSG,
		})
		return
	}
	if err := glo.Db.Model(&Project{}).Where("id = ?", reqData.ID).Delete(&Project{}).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "删除失败:" + fmt.Sprintf("%s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "删除成功",
	})
	return
}

func ProjectMsg(ctx *gin.Context) {
	type retData struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	var (
		res     []retData
		querSet []Project
	)

	if err := glo.Db.Model(&Project{}).Find(&querSet).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": e.ERROR_MSG + fmt.Sprintf("%s", err),
			"data":    []retData{},
		})
		return
	}
	for _, r := range querSet {
		res = append(res, retData{
			ID:   r.ID,
			Name: r.Name,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": e.SUCCESS_MSG,
		"data":    res,
	})
	return
}
