package sys

import (
	"fmt"
	"mos/src/glo"
	"mos/src/glo/comfunc"
	"mos/src/glo/encrypt"
	"net/http"
	"strconv"

	"mos/src/pkg/e"
	"mos/src/pkg/setting"
	"mos/src/pkg/util"
	"mos/src/server/route/jenkins"
	"mos/src/server/route/project"
	"mos/src/server/route/ticket"

	"github.com/gin-gonic/gin"
)

// ////////////////////////////////////////////////////////////////////////////////////////////////////////
// 路由函数
// ////////////////////////////////////////////////////////////////////////////////////////////////////////

// InitTable 初始化表结构
func InitTable(ctx *gin.Context) {
	if !glo.Db.HasTable(&jenkins.JenkinsJob{}) {
		if err := glo.Db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&jenkins.JenkinsJob{}).Error; err != nil {
			panic(err)
		}
	}
	if !glo.Db.HasTable(&User{}) {
		if err := glo.Db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
	if !glo.Db.HasTable(&SystemGroup{}) {
		if err := glo.Db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&SystemGroup{}).Error; err != nil {
			panic(err)
		}
	}
	if !glo.Db.HasTable(&Permission{}) {
		if err := glo.Db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Permission{}).Error; err != nil {
			panic(err)
		}
	}
	if !glo.Db.HasTable(&ticket.Ticket{}) {
		if err := glo.Db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&ticket.Ticket{}).Error; err != nil {
			panic(err)
		}
	}
	if !glo.Db.HasTable(&ticket.TicketHistory{}) {
		if err := glo.Db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&ticket.TicketHistory{}).Error; err != nil {
			panic(err)
		}
	}
	if !glo.Db.HasTable(&ticket.TicketType{}) {
		if err := glo.Db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&ticket.TicketType{}).Error; err != nil {
			panic(err)
		}
	}
	if !glo.Db.HasTable(&ticket.TicketSource{}) {
		if err := glo.Db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&ticket.TicketSource{}).Error; err != nil {
			panic(err)
		}
	}
	if !glo.Db.HasTable(&ticket.TicketLevel{}) {
		if err := glo.Db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&ticket.TicketLevel{}).Error; err != nil {
			panic(err)
		}
	}
	if !glo.Db.HasTable(&project.Project{}) {
		if err := glo.Db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&project.Project{}).Error; err != nil {
			panic(err)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "初始化表成功",
	})
}

// UserAdd 添加用户信息
func UserAdd(ctx *gin.Context) {
	var (
		user UserPostForm
		u    User
	)
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.SUCCESS,
			"message": "提交用户信息错误, " + err.Error(),
		})
		return
	}
	if err := u.findUserInfo(user.UserName); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.NOCONTENT,
			"message": "当前用户已存在, " + err.Error(),
		})
		return
	}
	// 设置默认密码
	if user.Password == `` {
		user.Password = comfunc.DefaultPassword
	}
	encryptPassword, _ := encrypt.AesEncryptString([]byte(user.Password), []byte(glo.Config.MosAPI.EncryptKey))
	// 映射POST数据到用户结构体
	insertData := &User{
		UserName: user.UserName,
		NickName: user.NickName,
		Password: encryptPassword,
		Email:    user.Email,
		Phone:    user.Phone,
	}
	if err := glo.Db.Create(&insertData).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "未知错误, " + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "添加用户成功",
	})
	return
}

// UserUpdate 更新用户信息
func UserUpdate(ctx *gin.Context) {
	var (
		user UserPostForm
		// u    User
	)
	userQuery := glo.Db.Table("User")
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": e.PARAM_ERROR_MSG + err.Error(),
		})
		return
	}
	// 映射POST数据到用户结构体
	insertData := &User{
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Phone:    user.Phone,
	}
	// 设置默认密码
	if user.Password != `` {
		encryptPassword, _ := encrypt.AesEncryptString([]byte(user.Password), []byte(glo.Config.MosAPI.EncryptKey))
		insertData.Password = encryptPassword
	}
	if err := userQuery.Where("id = ?", user.ID).Updates(&insertData).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": e.ERROR_MSG + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "更新用户信息成功",
	})
	return
}

// UserList 获取用户列表
func UserList(ctx *gin.Context) {
	// 初始化请求参数变量
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	page, _ := strconv.Atoi(ctx.Query("current_page"))
	pageSize = comfunc.GetDefaultPageSize(pageSize)
	page = comfunc.GetDefaultPage(page)
	// init params
	querySet := make([]User, 0)
	res := make([]UserInfo, 0)
	userQueryDb := glo.Db.Set("gorm:auto_preload", true)
	var (
		username string
		total    int
	)
	// 设置查询用户名
	username = ctx.Query("username")
	if username != `` {
		// 根据username模糊查询，可以将gorm链接添加条件后，赋值覆盖自身，得到不定条件的链式查询效果
		userQueryDb = userQueryDb.Where("username LIKE ?", fmt.Sprintf("%%%s%%", username))
	}
	if err := userQueryDb.Offset((page - 1) * pageSize).Limit(pageSize).Order("id desc").Find(&querySet).Count(&total).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "get user query error, " + err.Error(),
		})
		return
	}
	for _, r := range querySet {
		var v UserInfo
		// 校验密码是否为空，并解密，测试加解密代码准确定，生产环境注释以下if判断
		// if len(r.Password) != 0 {
		// 	decPassword, _ := encrypt.AesDecryptString(r.Password, []byte(glo.Config.MosAPI.EncryptKey))
		// 	v.Password = string(decPassword[:])
		// } else {
		// 	v.Password = ""
		// }
		v.ID = r.ID
		v.UserName = r.UserName
		v.NickName = r.NickName
		v.Email = r.Email
		v.Phone = r.Phone
		v.CreatedAt = comfunc.FormatTs(r.CreatedAt.Unix())
		for _, i := range r.SystemGroups {
			v.Groups = append(v.Groups, GroupInfSim{ID: i.ID, NickName: i.NickName, GroupName: i.GroupName})
		}
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

// UserDelete 删除用户
func UserDelete(ctx *gin.Context) {
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
	if err = glo.Db.Table("User").Where("id = ?", reqData.ID).Unscoped().Delete(User{}).Error; err != nil {
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

// UserLogin 用户登录
func UserLogin(ctx *gin.Context) {
	// UserLoginRet 用户登录后返回信息
	type UserLoginRet struct {
		Token string `json:"token"`
		Name  string `json:"name"`
	}
	var (
		user     UserPostForm
		u        User
		token    string
		loginRet UserLoginRet
	)
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": "用户登录失败",
			"data":    "",
		})
		return
	}
	// 校验密码
	encryptPassword, _ := encrypt.AesEncryptString([]byte(user.Password), []byte(glo.Config.MosAPI.EncryptKey))
	_, err = u.checkUserPassword(user.UserName, encryptPassword)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": fmt.Sprintf("用户登录失败: %s", err),
			"data":    "",
		})
		return
	}
	// JWT middleware生成token
	token, err = util.GenerateToken(user.UserName)

	// // 生成token, 旧方式
	// token = comfunc.EncryptToken(user.UserName, time.Now().Unix(), glo.Config.MosAPI.EncryptKey)
	// // 设置token缓存
	// rdsClient := glo.RdsDb.Get()
	// redisKey := fmt.Sprintf("auth_account_%s", token)
	// _, err = rdsClient.Do("set", redisKey, 1, "EX", glo.Config.MosAPI.Redis.Expried)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": err,
			"data":    "",
		})
	} else {
		loginRet.Token = token
		loginRet.Name = user.UserName
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.SUCCESS,
			"message": "login success",
			"data":    loginRet,
		})
	}
}

// UserLogout 用户登录
func UserLogout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "logout success",
	})
}

// AccountInfo 用户信息
func AccountInfo(ctx *gin.Context) {
	accountMsg := AccountMsg{
		Name:     "nil",
		NickName: "nil",
		Roles:    []string{},
		Perms:    []string{},
		Avatar:   setting.Avatar,
	}
	var (
		u User
	)

	token := ctx.GetHeader("X-Token")
	if token == `` {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": comfunc.TokenInvaild,
			"data":    accountMsg,
		})
		return
	}
	claims, err := util.ParseToken(token)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": err,
			"data":    accountMsg,
		})
		return
	}
	accountMsg.Name = claims.Username
	// 获取用户组信息
	if err = glo.Db.Set("gorm:auto_preload", true).Model(&User{}).Where("username = ?", accountMsg.Name).First(&u).Error; err == nil {
		for _, i := range u.SystemGroups {
			accountMsg.Roles = append(accountMsg.Roles, i.GroupName)
		}
		accountMsg.Perms, _ = u.getPermission()
		accountMsg.NickName = u.NickName
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "success",
		"data":    accountMsg,
	})
	return
}

// UserMsg 获取简单的全部用户信息
func UserMsg(ctx *gin.Context) {
	var userArr []UserMsgSim
	userQueryDb := glo.Db
	queryColoumn := []string{"id", "username", "nick_name"}
	if err := userQueryDb.Table("User").Select(queryColoumn).Find(&userArr).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "查询失败:" + err.Error(),
			"data":    "",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": e.SUCCESS_MSG,
		"data":    userArr,
	})
	return
}
