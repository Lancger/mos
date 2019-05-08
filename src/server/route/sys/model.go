package sys

import (
	"errors"
	"mos/src/glo"

	"github.com/jinzhu/gorm"
)

// User 用户表结构体
type User struct {
	// ID       int    `gorm:"type:int(10);primary_key;column:id;AUTO_INCREMENT;"`
	UserName     string        `gorm:"type:varchar(32);not null;column:username"`
	NickName     string        `gorm:"type:varchar(32);not null;column:nick_name"`
	Password     string        `gorm:"type:varchar(64);column:password"`
	Email        string        `gorm:"type:varchar(64);column:email"`
	Phone        string        `gorm:"type:varchar(16);column:phone"`
	SystemGroups []SystemGroup `gorm:"many2many:user_group"`
	gorm.Model
	// CreatedTime  int64         `gorm:"column:create_time"`
}

// UserInfo 用户信息返回
type UserInfo struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	// Password    string `json:"password"`
	Groups    []GroupInfSim `json:"groups"`
	CreatedAt string        `json:"create_time"`
}
type GroupInfSim struct {
	ID        uint   `json:"id"`
	GroupName string `json:"group_name"`
	NickName  string `json:"nick_name"`
}

// 模糊查询用户名结构体
type userGetData struct {
	PageSize int    `json:"page_size"`
	Page     int    `json:"page"`
	Username string `json:"username"`
}

// UserPostForm 接收请求POST数据结构体
type UserPostForm struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	NickName string `json:"nick_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// SystemGroup 用户组表结构体
type SystemGroup struct {
	// ID          int    `gorm:"type:int;primary_key;column:id;AUTO_INCREMENT"`
	GroupName string `gorm:"type:varchar(32);not null;column:group_name"`
	NickName  string `gorm:"type:varchar(32);not null;column:nick_name"`
	Users     []User `gorm:"many2many:user_group"`
	gorm.Model
	// CreatedTime int64  `gorm:"column:create_time"`
}

// UserGroupMap 用户组与用户关系中间表结构体
type UserGroupMap struct {
	UserID  uint `gorm:"type:int(10);not null;column:user_id"`
	GroupID uint `gorm:"type:int(10);not null;column:system_group_id"`
}

// GroupInfo 用户信息返回
type GroupInfo struct {
	ID        uint   `json:"id"`
	GroupName string `json:"group_name"`
	NickName  string `json:"nick_name"`
	CreatedAt string `json:"create_time"`
}

type groupGetData struct {
	PageSize  int    `json:"page_size"`
	Page      int    `json:"page"`
	GroupName string `json:"group_name"`
}

// GroupPostForm 接收请求POST数据结构体
type GroupPostForm struct {
	GroupName string `json:"group_name"`
	NickName  string `json:"nick_name"`
	UserIds   []uint `json:"user_selected"`
}

// TableName User 用户组表名
func (SystemGroup) TableName() string {
	return "SystemGroup"
}

// TableName User 用户表名
func (UserGroupMap) TableName() string {
	return "user_group"
}

// AccountMsg 账号信息
type AccountMsg struct {
	Name   string   `json:"name"`
	Roles  []string `json:"roles"`
	Perms  []string `json:"perms"`
	Avatar string   `json:"avatar"`
}

// UserMsgSim 用户信息
type UserMsgSim struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	NickName string `json:"nick_name"`
}

// findGroupInfo 查找是否已有用户组信息
func (g SystemGroup) findGroupInfo(groupName string) (err error) {
	var count int
	err = glo.Db.Model(&SystemGroup{}).Where(&SystemGroup{GroupName: groupName}).Count(&count).Error
	if err != nil {
		return
	}

	if count > 0 {
		err = errors.New("group is existed")
	}
	return
}

// findGroupNickName 查找是否已有用户组信息
func (g SystemGroup) findGroupNickName(nickName string) (err error) {
	var count int
	err = glo.Db.Model(&SystemGroup{}).Where(&SystemGroup{NickName: nickName}).Count(&count).Error
	if err != nil {
		return
	}

	if count > 0 {
		err = errors.New("group is existed")
	}
	return
}

// TableName User 用户表名
func (User) TableName() string {
	return "User"
}

// findUserInfo 查找是否已有用户信息
func (u User) findUserInfo(username string) (err error) {
	var count int
	err = glo.Db.Model(&User{}).Where(&User{UserName: username}).Count(&count).Error
	if err != nil {
		return
	}

	if count > 0 {
		err = errors.New("user is existed")
	}
	return
}
func (u User) checkUserExists(username string) (status bool, err error) {
	var count int
	err = glo.Db.Model(&User{}).Where(&User{UserName: username}).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count <= 0 {
		err = errors.New("count less than 0")
		return false, err
	}
	return true, nil
}

func (u User) checkUserPassword(username string, password string) (status bool, err error) {
	var (
		userMsg User
	)
	err = glo.Db.Model(&User{}).Where(&User{UserName: username}).First(&userMsg).Error
	if err != nil {
		return false, err
	}
	if password != userMsg.Password {
		err = errors.New("password error")
		return false, err
	}
	return true, nil
}

// findUserId 查找是否已有用户信息
func (u User) findUserID(id uint) (err error) {
	var count int
	u.ID = id
	err = glo.Db.Model(&User{}).Where(&u).Count(&count).Error
	if err != nil {
		return
	}

	if count > 0 {
		err = errors.New("user is existed")
	}
	return
}
