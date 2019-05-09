package sys

import (
	"errors"
	"log"
	"mos/src/glo"

	"github.com/jinzhu/gorm"
)

// User 用户表结构体
type User struct {
	// ID       int    `gorm:"type:int(10);primary_key;column:id;AUTO_INCREMENT;"`
	UserName     string        `gorm:"type:varchar(32);unique;not null;column:username"`
	NickName     string        `gorm:"type:varchar(32);not null;column:nick_name"`
	Password     string        `gorm:"type:varchar(64);column:password"`
	Email        string        `gorm:"type:varchar(64);column:email"`
	Phone        string        `gorm:"type:varchar(16);column:phone"`
	SystemGroups []SystemGroup `gorm:"many2many:user_group"`
	gorm.Model
}

// Permission 用户权限
type Permission struct {
	Name         string        `gorm:"type:varchar(32);unique;not null;column:name"`
	NickName     string        `gorm:"type:varchar(32);not null;column:nick_name"`
	Type         string        `gorm:"type:varchar(32);not null;column:type"`
	SystemGroups []SystemGroup `gorm:"many2many:group_permission"`
	gorm.Model
}

// SystemGroup 用户组表结构体
type SystemGroup struct {
	// ID          int    `gorm:"type:int;primary_key;column:id;AUTO_INCREMENT"`
	GroupName   string       `gorm:"type:varchar(32);unique;not null;column:group_name"`
	NickName    string       `gorm:"type:varchar(32);not null;column:nick_name"`
	Users       []User       `gorm:"many2many:user_group"`
	Permissions []Permission `gorm:"many2many:group_permission"`
	gorm.Model
}

// UserGroupMap 用户组与用户关系中间表结构体
type UserGroupMap struct {
	UserID  uint `gorm:"type:int(10);not null;column:user_id"`
	GroupID uint `gorm:"type:int(10);not null;column:system_group_id"`
}

// GroupPermMap 用户组权限表关联
type GroupPermMap struct {
	PermissionID uint `gorm:"type:int(10);not null;column:permission_id"`
	GroupID      uint `gorm:"type:int(10);not null;column:system_group_id"`
}

// UserInfo 用户信息返回
type UserInfo struct {
	ID        uint          `json:"id"`
	UserName  string        `json:"username"`
	NickName  string        `json:"nick_name"`
	Email     string        `json:"email"`
	Phone     string        `json:"phone"`
	Groups    []GroupInfSim `json:"groups"`
	CreatedAt string        `json:"create_time"`
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

// GroupInfSim 简单用户组信息
type GroupInfSim struct {
	ID        uint   `json:"id"`
	GroupName string `json:"group_name"`
	NickName  string `json:"nick_name"`
}

// GroupInfo 用户信息返回
type GroupInfo struct {
	ID        uint   `json:"id"`
	GroupName string `json:"group_name"`
	NickName  string `json:"nick_name"`
	CreatedAt string `json:"create_time"`
}

// PermissionInfo 用户权限
type PermissionInfo struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	NickName string `json:"nick_name"`
	Type     string `json:"type"`
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

// TableName User 用户表名
func (User) TableName() string {
	return "User"
}

// TableName User 用户组表名
func (SystemGroup) TableName() string {
	return "SystemGroup"
}

// TableName User 表名
func (UserGroupMap) TableName() string {
	return "user_group"
}

// TableName Permission 权限表名
func (Permission) TableName() string {
	return "Permission"
}

// TableName GroupPermMap 权限中间表表名
func (GroupPermMap) TableName() string {
	return "group_permission"
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

// getPermission 获取用户权限
func (u User) getPermission() (permissionArr []string, err error) {
	var (
		groupID  []uint
		groupArr []SystemGroup
	)

	if err = glo.Db.Set("gorm:auto_preload", true).Model(&User{}).First(&u).Error; err == nil {
		for _, i := range u.SystemGroups {
			groupID = append(groupID, i.ID)
			log.Printf("%d", i.ID)
		}

		if err = glo.Db.Set("gorm:auto_preload", true).Model(&SystemGroup{}).Where("id in (?)", groupID).Find(&groupArr).Error; err == nil {
			for _, g := range groupArr {
				for _, p := range g.Permissions {
					// log.Printf(p.NickName)
					permissionArr = append(permissionArr, p.Name)
				}
			}
		}
	} else {
		permissionArr = []string{}
	}
	return
}
