package sys

import (
	"fmt"
	"mos/src/glo"
	"mos/src/glo/comfunc"
	"mos/src/pkg/e"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ////////////////////////////////////////////////////////////////////////////////////////////////////////
// 路由函数
// ////////////////////////////////////////////////////////////////////////////////////////////////////////

// PermissionList 获取所有权限
func PermissionList(ctx *gin.Context) {
	type permInfo struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		Type     string `json:"type"`
		NickName string `json:"nick_name"`
	}
	var (
		g        []Permission
		permList []permInfo
	)
	if err := glo.Db.Set("gorm:auto_preload", true).Model(&Permission{}).Find(&g).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "获取权限信息失败",
			"data":    []string{},
		})
		return
	}
	for _, p := range g {
		permList = append(permList, permInfo{ID: p.ID, Name: p.Name, NickName: p.NickName, Type: p.Type})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "获取权限信息成功",
		"data":    permList,
	})
	return
}

// GroupPermission 根据用户组ID获取权限ID数组
func GroupPermission(ctx *gin.Context) {
	var (
		g SystemGroup
	)
	permIDList := make([]uint, 0)
	id := ctx.Query("id")
	if id == `` {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": "请求参数错误",
			"data":    permIDList,
		})
		return
	}
	if err := glo.Db.Set("gorm:auto_preload", true).Model(&SystemGroup{}).Where("id = ?", id).First(&g).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "获取权限信息失败：" + fmt.Sprintf("%s", err),
			"data":    permIDList,
		})
		return
	}
	for _, v := range g.Permissions {
		permIDList = append(permIDList, v.ID)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "获取权限信息成功",
		"data":    permIDList,
	})
	return
}

// PermissionMsg 获取权限列表
func PermissionMsg(ctx *gin.Context) {
	// 初始化请求参数变量
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	page, _ := strconv.Atoi(ctx.Query("current_page"))
	pageSize = comfunc.GetDefaultPageSize(pageSize)
	page = comfunc.GetDefaultPage(page)
	// init params
	querySet := make([]Permission, 0)
	res := make([]PermissionInfo, 0)
	queryDb := glo.Db
	var (
		nickName string
		total    int
	)
	// 设置查询权限
	nickName = ctx.Query("search")
	if nickName != `` {
		// 根据nickName模糊查询，可以将gorm链接添加条件后，赋值覆盖自身，得到不定条件的链式查询效果
		queryDb = queryDb.Where("nick_name LIKE ?", fmt.Sprintf("%%%s%%", nickName))
	}
	if err := queryDb.Offset((page - 1) * pageSize).Limit(pageSize).Order("id desc").Find(&querySet).Count(&total).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "get perm query error, " + err.Error(),
			"data":    res,
		})
		return
	}
	for _, r := range querySet {
		var v PermissionInfo
		v.ID = r.ID
		v.Name = r.Name
		v.NickName = r.NickName
		v.Type = r.Type
		res = append(res, v)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":       20000,
		"message":    "success",
		"data":       res,
		"total":      total,
		"total_page": comfunc.FlorPageInt(pageSize, total),
		"page_size":  pageSize,
		"page":       page,
	})
	return
}

// PermissionAdd 添加权限
func PermissionAdd(ctx *gin.Context) {
	type reqPostData struct {
		Type     string `json:"type"`
		Name     string `json:"name"`
		NickName string `json:"nick_name"`
	}
	var (
		reqData reqPostData
	)
	if err := ctx.BindJSON(&reqData); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": e.PARAM_ERROR_MSG,
		})
		return
	}
	p := Permission{
		Name:     reqData.Name,
		NickName: reqData.NickName,
		Type:     reqData.Type,
	}
	if err := glo.Db.Create(&p).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "添加权限失败:" + fmt.Sprintf("%s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "添加权限成功",
	})
	return
}

// PermissionDelete 删除权限
func PermissionDelete(ctx *gin.Context) {
	type requestPost struct {
		ID int `json:"id"`
	}
	var reqData requestPost
	err := ctx.BindJSON(&reqData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "请求参数错误" + err.Error(),
		})
		return
	}
	glo.Db.Model(&GroupPermMap{}).Where("permission_id = ?", reqData.ID).Unscoped().Delete(GroupPermMap{})
	// Unscoped 永久删除，否则gin只会软删除
	if err = glo.Db.Model(&Permission{}).Where("id = ?", reqData.ID).Unscoped().Delete(Permission{}).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "删除失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "删除成功",
	})
	return
}

// PermissionUpdate 权限信息更新
func PermissionUpdate(ctx *gin.Context) {
	type requestPost struct {
		ID       int    `json:"id"`
		Type     string `json:"type"`
		Name     string `json:"name"`
		NickName string `json:"nick_name"`
	}
	var reqData requestPost
	err := ctx.BindJSON(&reqData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "请求参数错误" + err.Error(),
		})
		return
	}
	updateData := Permission{
		Type:     reqData.Type,
		Name:     reqData.Name,
		NickName: reqData.NickName,
	}
	if err = glo.Db.Model(&Permission{}).Where("id = ?", reqData.ID).Update(updateData).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "更新权限信息失败" + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "更新权限信息成功",
	})
	return
}
