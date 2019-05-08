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

// GroupAdd 添加用户组信息
func GroupAdd(ctx *gin.Context) {
	var (
		group GroupPostForm
		g     SystemGroup
	)

	err := ctx.BindJSON(&group)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": "用户组提交信息错误, " + err.Error(),
		})
		return
	}
	if err := g.findGroupInfo(group.GroupName); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.NOCONTENT,
			"message": "用户组已存在, " + err.Error(),
		})
		return
	}
	// 映射POST数据到用户结构体
	insertData := &SystemGroup{
		GroupName: group.GroupName,
		NickName:  group.NickName,
	}
	if err := glo.Db.Create(&insertData).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "未知错误, " + err.Error(),
		})
		return
	}
	// 用户组与用户进行关联
	if len(group.UserIds) > 0 {
		var s SystemGroup
		if err = glo.Db.Model(&SystemGroup{}).Where("group_name = ?", group.GroupName).First(&s).Error; err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    e.WARNNING,
				"message": "添加用户组成功，但管理用户失败",
			})
			return
		}
		var mapData UserGroupMap
		for _, v := range group.UserIds {
			mapData = UserGroupMap{
				UserID:  v,
				GroupID: s.ID,
			}
			glo.Db.Create(&mapData)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "添加用户组成功",
	})
	return
}

// GroupUpdate 更新用户组信息
func GroupUpdate(ctx *gin.Context) {
	type postData struct {
		ID           uint   `json:"id"`
		NickName     string `json:"nick_name"`
		UserSelected []uint `json:"user_selected"`
	}
	var (
		reqData postData
	)

	err := ctx.BindJSON(&reqData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": "用户组提交信息错误, " + err.Error(),
		})
		return
	}

	if err := glo.Db.Model(&SystemGroup{}).Where("id = ?", reqData.ID).Update(SystemGroup{NickName: reqData.NickName}).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "未知错误, " + err.Error(),
		})
		return
	}
	glo.Db.Model(&UserGroupMap{}).Where("system_group_id = ?", reqData.ID).Unscoped().Delete(&UserGroupMap{})
	// 用户组与用户进行关联
	if len(reqData.UserSelected) > 0 {
		var mapData UserGroupMap
		for _, v := range reqData.UserSelected {
			mapData = UserGroupMap{
				UserID:  v,
				GroupID: reqData.ID,
			}
			glo.Db.Create(&mapData)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "修改用户组成功",
	})
	return
}

// GroupList 获取用户组列表
func GroupList(ctx *gin.Context) {
	// 初始化请求参数变量
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize = comfunc.GetDefaultPageSize(pageSize)
	page = comfunc.GetDefaultPage(page)
	// init params
	querySet := make([]SystemGroup, 0)
	res := make([]GroupInfo, 0)
	groupQuery := glo.Db.Model(&SystemGroup{})
	var (
		groupname string
		total     int
	)
	// 设置查询用户名
	groupname = ctx.Query("group_name")
	if groupname != `` {
		// 根据groupname模糊查询，可以将gorm链接添加条件后，赋值覆盖自身，得到不定条件的链式查询效果
		groupQuery = groupQuery.Where("group_name LIKE ?", fmt.Sprintf("%%%s%%", groupname))
	}

	if err := groupQuery.Limit(pageSize).Offset((page - 1) * pageSize).Order("id desc").Find(&querySet).Count(&total).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "查询用户组失败 " + err.Error(),
		})
		return
	}
	for _, r := range querySet {
		var g GroupInfo
		g.ID = r.ID
		g.GroupName = r.GroupName
		g.NickName = r.NickName
		g.CreatedAt = comfunc.FormatTs(r.CreatedAt.Unix())
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

// GroupDelete 删除用户组
func GroupDelete(ctx *gin.Context) {
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
	// Unscoped 永久删除，否则gin只会软删除
	if err = glo.Db.Table("SystemGroup").Where("id = ?", reqData.ID).Unscoped().Delete(SystemGroup{}).Error; err != nil {
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

// GroupMsg 根据用户组ID获取所属用户ID列表
func GroupMsg(ctx *gin.Context) {
	id := ctx.Query("id")
	userIDArr := []uint{}

	if id == `` {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "请求参数错误",
			"data":    userIDArr,
		})
		return
	}
	var m []UserGroupMap
	if err := glo.Db.Model(&UserGroupMap{}).Where("system_group_id = ?", id).Find(&m).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "查询失败",
			"data":    userIDArr,
		})
		return
	}
	for _, r := range m {
		userIDArr = append(userIDArr, r.UserID)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "获取成功",
		"data":    userIDArr,
	})
	return
}
