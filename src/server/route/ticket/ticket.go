package ticket

import (
	"fmt"
	"mos/src/glo"
	"mos/src/glo/comfunc"
	"mos/src/pkg/e"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const ConstTicketNew string = "新建工单"
const ConstTicketReject string = "驳回"
const ConstTicketHold string = "暂存"
const ConstTicketFinish string = "处理完毕"
const ConstTicketClose string = "结单"
const ConstTicketDealing string = "处理中"

// TicketSourceAdd 添加工单类型
func TicketSourceAdd(ctx *gin.Context) {
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
	t := TicketSource{
		Name: r.Name,
	}
	if err := glo.Db.Create(&t).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "添加工单来源失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "添加工单来源成功",
	})
	return
}

// TicketSourceList 工单来源列表
func TicketSourceList(ctx *gin.Context) {
	// 初始化请求参数变量
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize = comfunc.GetDefaultPageSize(pageSize)
	page = comfunc.GetDefaultPage(page)
	// init params
	querySet := make([]TicketSource, 0)
	res := make([]TicketSourceInfo, 0)
	dbQuery := glo.Db.Model(&TicketSource{})
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
		var g TicketSourceInfo
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

// TicketSourceMsg 获取所有工单来源
func TicketSourceMsg(ctx *gin.Context) {
	type retJson struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	var (
		retData    []retJson
		sourceList []TicketSource
	)

	if err := glo.Db.Model(&TicketSource{}).Find(&sourceList).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "查询失败:" + err.Error(),
			"data":    []retJson{},
		})
		return
	}
	var sig retJson
	for _, r := range sourceList {
		sig.ID = r.ID
		sig.Name = r.Name
	}
	retData = append(retData, sig)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": e.SUCCESS_MSG,
		"data":    retData,
	})
	return

}

// TicketSourceUpdate 更新工单类型
func TicketSourceUpdate(ctx *gin.Context) {
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
	if err := glo.Db.Model(&TicketSource{}).Where("id = ?", reqData.ID).Update(&TicketSource{Name: reqData.Name}).Error; err != nil {
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

// TicketSourceUpdate 更新工单类型
func TicketSourceDelete(ctx *gin.Context) {
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
	if err := glo.Db.Model(&TicketSource{}).Where("id = ?", reqData.ID).Delete(&TicketSource{}).Error; err != nil {
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

// TicketTypeAdd 添加工单类型
func TicketTypeAdd(ctx *gin.Context) {
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
	t := TicketType{
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

// TicketTypeList 工单类型列表
func TicketTypeList(ctx *gin.Context) {
	// 初始化请求参数变量
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize = comfunc.GetDefaultPageSize(pageSize)
	page = comfunc.GetDefaultPage(page)
	// init params
	querySet := make([]TicketType, 0)
	res := make([]TicketTypeInfo, 0)
	dbQuery := glo.Db.Model(&TicketType{})
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
		var g TicketTypeInfo
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

// TicketTypeMsg 获取所有工单类型
func TicketTypeMsg(ctx *gin.Context) {
	type retJson struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	var (
		retData  []retJson
		typeList []TicketType
	)

	if err := glo.Db.Model(&TicketType{}).Find(&typeList).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "查询失败:" + err.Error(),
			"data":    []retJson{},
		})
		return
	}
	var sig retJson
	for _, r := range typeList {
		sig.ID = r.ID
		sig.Name = r.Name
	}
	retData = append(retData, sig)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": e.SUCCESS_MSG,
		"data":    retData,
	})
	return

}

// TicketTypeUpdate 更新工单类型
func TicketTypeUpdate(ctx *gin.Context) {
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
	if err := glo.Db.Model(&TicketType{}).Where("id = ?", reqData.ID).Update(&TicketType{Name: reqData.Name}).Error; err != nil {
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

// TicketTypeUpdate 更新工单类型
func TicketTypeDelete(ctx *gin.Context) {
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
	if err := glo.Db.Model(&TicketType{}).Where("id = ?", reqData.ID).Delete(&TicketType{}).Error; err != nil {
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

// TicketLevelAdd 添加工单等级
func TicketLevelAdd(ctx *gin.Context) {
	type reqPostData struct {
		Name  string `json:"name"`
		Level uint   `json:"level"`
		Time  uint   `json:"time"`
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
	t := TicketLevel{
		Name:  r.Name,
		Level: r.Level,
		Time:  r.Time,
	}
	if err := glo.Db.Create(&t).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "添加工单等级失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "添加工单等级成功",
	})
	return
}

// TicketLevelDelete 更新工单等级
func TicketLevelDelete(ctx *gin.Context) {
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
	if err := glo.Db.Model(&TicketLevel{}).Where("id = ?", reqData.ID).Delete(&TicketLevel{}).Error; err != nil {
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

// TicketLevelList 工单等级列表
func TicketLevelList(ctx *gin.Context) {
	// 初始化请求参数变量
	// init params
	querySet := make([]TicketLevel, 0)
	res := make([]TicketLevelInfo, 0)
	dbQuery := glo.Db.Model(&TicketLevel{})

	if err := dbQuery.Find(&querySet).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "查询失败 " + err.Error(),
		})
		return
	}
	for _, r := range querySet {
		var g TicketLevelInfo
		g.ID = r.ID
		g.Name = r.Name
		g.Time = r.Time
		g.Level = r.Level
		g.CreateTime = comfunc.FormatTs(r.CreatedAt.Unix())
		res = append(res, g)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "success",
		"data":    res,
	})
	return
}

// TicketLevelUpdate 更新工单类型
func TicketLevelUpdate(ctx *gin.Context) {
	// type reqPostData struct {
	// 	ID   uint   `json:"id"`
	// 	Name string `json:"name"`
	// 	level string `json:"level"`
	// 	Time uint `json:"time"`
	// }
	// var reqData reqPostData
	var reqData TicketLevel
	if err := ctx.BindJSON(&reqData); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": e.PARAM_ERROR_MSG,
		})
		return
	}

	if err := glo.Db.Model(&TicketLevel{}).Where("id = ?", reqData.ID).Updates(&reqData).Error; err != nil {
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

// TicketList 工单列表
func TicketList(ctx *gin.Context) {
	type reqData struct {
		Page      int      `json:"current_page"`
		PageSize  int      `json:"page_size"`
		Search    string   `json:"search"`
		Number    string   `json:"ticket_number"`
		Dealer    string   `json:"dealer"`
		Author    string   `json:"author"`
		Project   []string `json:"project"`
		Type      []string `json:"type"`
		Stage     []string `json:"stage"`
		UnFinish  uint     `json:"unfinish"`
		StartTime int64    `json:"start_time"`
		EndTime   int64    `json:"end_time"`
	}
	var (
		req      reqData
		t        TicketInfo
		pageSize int
		page     int
		total    int
	)

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":       e.PARAM_ERROR,
			"message":    e.PARAM_ERROR_MSG,
			"data":       []TicketInfo{},
			"total":      0,
			"total_page": 0,
			"page_size":  0,
			"page":       0,
		})
		return
	}

	// 初始化请求参数变量
	pageSize = comfunc.GetDefaultPageSize(req.PageSize)
	page = comfunc.GetDefaultPage(req.Page)
	// init params
	querySet := make([]Ticket, 0)
	res := make([]TicketInfo, 0)
	queryDb := glo.Db.Set("gorm:auto_preload", true).Model(&Ticket{})
	if req.Author != `` {
		queryDb = queryDb.Where("author = ?", req.Author)
	}
	if len(req.Project) > 0 {
		queryDb = queryDb.Where("project IN (?)", req.Project)
	}
	if req.Search != `` {
		queryDb = queryDb.Where("content LIKE ?", fmt.Sprintf("%%%s%%", req.Search))
	}
	if req.Number != `` {
		queryDb = queryDb.Where("number = ?", req.Number)
	}
	if req.Dealer != `` {
		queryDb = queryDb.Where("dealer = ?", req.Dealer)
	}
	if len(req.Type) > 0 {
		queryDb = queryDb.Where("type IN (?)", req.Type)
	}
	if len(req.Stage) > 0 {
		queryDb = queryDb.Where("dealer IN (?)", req.Stage)
	}
	if req.UnFinish == 1 {
		queryDb = queryDb.Where("stage != ?", ConstTicketClose)
	}
	if req.StartTime > 0 {
		queryDb = queryDb.Where("created_at >= ?", comfunc.FormatTs(req.StartTime/1e3))
	}
	if req.EndTime > 0 {
		queryDb = queryDb.Where("created_at <= ?", comfunc.FormatTs(req.EndTime/1e3))
	}
	if err := queryDb.Offset((page - 1) * pageSize).Limit(pageSize).Order("id desc").Find(&querySet).Count(&total).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "get user query error, " + err.Error(),
		})
		return
	}
	for _, r := range querySet {
		s := []TicketHistoryInfo{}
		for _, h := range r.TicketHistory {
			s = append(s, TicketHistoryInfo{
				ID:         h.ID,
				Action:     h.Action,
				Content:    h.Content,
				From:       h.From,
				To:         h.To,
				CreateTime: comfunc.FormatTs(h.CreatedAt.Unix()),
			})
		}

		t = TicketInfo{
			ID:         r.ID,
			Number:     r.Number,
			Title:      r.Title,
			Level:      r.Level,
			Type:       r.Type,
			Stage:      r.Stage,
			Source:     r.Source,
			Improve:    r.Improve,
			Content:    r.Content,
			Solution:   r.Solution,
			Dealer:     r.Dealer,
			PreDealer:  r.PreDealer,
			Author:     r.Author,
			Project:    r.Project,
			Histiory:   s,
			CreateTime: comfunc.FormatTs(r.CreatedAt.Unix()),
		}
		res = append(res, t)
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

// TicketAdd 添加工单
func TicketAdd(ctx *gin.Context) {
	type reqData struct {
		Title     string `json:"title"`
		Level     uint   `json:"level"`
		Type      string `json:"type"`
		Stage     string `json:"stage"`
		Source    string `json:"source"`
		Improve   string `json:"improve"`
		Content   string `json:"content"`
		Solution  string `json:"solution"`
		Dealer    string `json:"dealer"`
		PreDealer string `json:"pre_dealer"`
		Author    string `json:"author"`
		Project   string `json:"project"`
	}
	var (
		req reqData
		t   Ticket
	)
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": e.PARAM_ERROR_MSG,
		})
		return
	}
	t = Ticket{
		Number:    comfunc.FormatShortTs(time.Now().Unix()),
		Title:     req.Title,
		Level:     req.Level,
		Type:      req.Type,
		Stage:     req.Stage,
		Project:   req.Project,
		Source:    req.Source,
		Improve:   req.Improve,
		Content:   req.Content,
		Solution:  req.Solution,
		Dealer:    req.Dealer,
		PreDealer: req.PreDealer,
		Author:    req.Author,
	}
	if err := glo.Db.Model(&Ticket{}).Create(&t).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": e.ERROR_MSG,
			"debug":   "添加失败:" + fmt.Sprintf("%s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": e.SUCCESS_MSG,
	})
	return
}

// TicketUpdate 更新工单
func TicketUpdate(ctx *gin.Context) {
	type reqData struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Level     uint   `json:"level"`
		Type      string `json:"type"`
		Stage     string `json:"stage"`
		Source    string `json:"source"`
		Improve   string `json:"improve"`
		Content   string `json:"content"`
		Solution  string `json:"solution"`
		Dealer    string `json:"dealer"`
		PreDealer string `json:"pre_dealer"`
		Author    string `json:"author"`
		Project   string `json:"project"`
	}
	var (
		req reqData
	)
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": e.PARAM_ERROR_MSG,
		})
		return
	}
	t := Ticket{
		Title:    req.Title,
		Level:    req.Level,
		Type:     req.Type,
		Stage:    req.Stage,
		Source:   req.Source,
		Improve:  req.Improve,
		Content:  req.Content,
		Solution: req.Solution,
		Project:  req.Project,
	}
	if err := glo.Db.Model(&Ticket{}).Where("id = ?", req.ID).Updates(&t).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": e.ERROR_MSG,
			"debug":   "修改失败:" + fmt.Sprintf("%s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": e.SUCCESS_MSG,
	})
	return
}

// TicketDelete 删除工单
func TicketDelete(ctx *gin.Context) {
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
	if err = glo.Db.Model(&Ticket{}).Where("id = ?", reqData.ID).Unscoped().Delete(Ticket{}).Error; err != nil {
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

// TicketMsg 根据工单号获取工单信息
func TicketMsg(ctx *gin.Context) {
	var (
		t Ticket
	)

	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": e.PARAM_ERROR_MSG,
			"data":    []string{},
		})
		return
	}
	if err = glo.Db.Model(&Ticket{}).Where("number = ?", id).First(&t).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": e.PARAM_ERROR_MSG,
			"data":    []string{},
		})
		return
	}
	tInfo := TicketInfo{
		ID:         t.ID,
		Number:     t.Number,
		Title:      t.Title,
		Level:      t.Level,
		Type:       t.Type,
		Stage:      t.Stage,
		Source:     t.Source,
		Improve:    t.Improve,
		Content:    t.Content,
		Solution:   t.Solution,
		StartTime:  comfunc.FormatTs(t.StartTime),
		EndTime:    comfunc.FormatTs(t.EndTime),
		Dealer:     t.Dealer,
		PreDealer:  t.PreDealer,
		Author:     t.Author,
		Project:    t.Project,
		CreateTime: comfunc.FormatTs(t.CreatedAt.Unix()),
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": e.SUCCESS_MSG,
		"data":    tInfo,
	})
	return
}

// TicketSend 派发工单或转单
func TicketSend(ctx *gin.Context) {
	type reqData struct {
		ID      uint   `json:"id"`
		Action  string `json:"action"`
		Content string `json:"content"`
		From    string `json:"from"`
		To      string `json:"to"`
	}
	var (
		req reqData
		act string
		t   Ticket
	)
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": e.PARAM_ERROR_MSG,
		})
		return
	}
	if req.Action == "send" {
		act = "派发工单"
	} else if req.Action == "change" {
		act = "转发工单"
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "操作类型错误",
		})
		return
	}

	if err := glo.Db.Model(&Ticket{}).Where("id = ?", req.ID).First(&t).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": act + fmt.Sprintf("失败:%s", err),
		})
		return
	}
	// 工单完成后禁止转单
	if t.Stage == ConstTicketClose {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": act + fmt.Sprintf("失败:当前工单已完结"),
		})
		return
	}
	if err := glo.Db.Model(&Ticket{}).Where("id = ?", req.ID).Updates(Ticket{Dealer: req.To, PreDealer: t.Dealer}).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": act + fmt.Sprintf("失败:%s", err),
		})
		return
	}
	h := TicketHistory{
		TicketID: req.ID,
		Action:   act,
		Content:  req.Content,
		From:     req.From,
		To:       req.To,
	}
	if err := glo.Db.Model(&TicketHistory{}).Create(&h).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "添加失败:" + fmt.Sprintf("%s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": e.SUCCESS_MSG,
	})
	return
}

func TicketCtl(ctx *gin.Context) {
	type reqData struct {
		ID      uint   `json:"id"`
		Ctl     string `json:"ctl"`
		User    string `json:"user"`
		Content string `json:"content"`
	}

	var (
		req   reqData
		t     Ticket
		h     TicketHistory
		stage string
	)
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": e.PARAM_ERROR_MSG,
		})
		return
	}
	if err := glo.Db.Model(&Ticket{}).Where("id = ?", req.ID).Find(&t).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": e.ERROR_MSG,
		})
		return
	}
	if req.Ctl == "deal" {
		for _, c := range []string{ConstTicketClose, ConstTicketFinish} {
			if t.Stage == c {
				ctx.JSON(http.StatusOK, gin.H{
					"code":    e.ERROR,
					"message": "工单已处理完毕，禁止再处理, 请修改工单状态",
				})
				return
			}
		}
		stage = ConstTicketDealing
		h = TicketHistory{
			Action:  "开始处理",
			Content: "开始处理工单",
			From:    req.User,
			To:      req.User,
		}

	} else if req.Ctl == "mark" {
		if t.Stage == ConstTicketNew {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    e.ERROR,
				"message": "新建工单禁止添加记录，请先开始处理工单",
			})
			return
		}
		stage = t.Stage
		h = TicketHistory{
			Action:  "添加处理记录",
			Content: req.Content,
			From:    req.User,
			To:      req.User,
		}
	} else if req.Ctl == "finish" {
		if t.Stage == ConstTicketClose {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    e.ERROR,
				"message": "工单已结单",
			})
			return
		}
		stage = ConstTicketFinish
		h = TicketHistory{
			Action:  "处理完毕",
			Content: req.Content,
			From:    req.User,
			To:      req.User,
		}
	} else if req.Ctl == "close" {
		if t.Stage == ConstTicketClose {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    e.ERROR,
				"message": "工单已结单",
			})
			return
		}
		stage = ConstTicketClose
		h = TicketHistory{
			Action:  "结单",
			Content: req.Content,
			From:    req.User,
			To:      req.User,
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "ctl操作参数错误",
		})
		return
	}
	if err := glo.Db.Model(&Ticket{}).Where("id = ?", req.ID).Update(&Ticket{Stage: stage}).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": e.ERROR_MSG,
		})
		return
	}
	h.TicketID = req.ID
	if err := glo.Db.Model(&TicketHistory{}).Create(&h).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.WARNNING,
			"message": "警告:" + fmt.Sprintf("%s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": e.SUCCESS_MSG,
	})
	return
}
