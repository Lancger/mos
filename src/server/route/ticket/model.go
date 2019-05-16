package ticket

import (
	"github.com/jinzhu/gorm"
)

// 工单名
type Ticket struct {
	Number        string `gorm:"type:varchar(128);not null;column:number"`
	Title         string `gorm:"type:varchar(128);not null;column:title"`
	Level         uint   `gorm:"type:int(10);not null;column:level"`
	Type          string `gorm:"type:varchar(128);not null;column:type"`
	Stage         string `gorm:"type:varchar(128);not null;column:stage"`
	Source        string `gorm:"type:varchar(128);not null;column:source"`
	Improve       string `gorm:"type:varchar(128);not null;column:improve"`
	Content       string `gorm:"type:mediumtext;not null;column:content"`
	Solution      string `gorm:"type:mediumtext;not null;column:solution"`
	StartTime     int64  `gorm:"type:int(64);default null;column:start_time"`
	EndTime       int64  `gorm:"type:int(64);default null;column:end_time"`
	Dealer        string `gorm:"type:varchar(128);not null;column:dealer"`
	PreDealer     string `gorm:"type:varchar(128);not null;column:pre_ealer"`
	Author        string `gorm:"type:varchar(128);not null;column:author"`
	Project       string `gorm:"type:varchar(128);not null;column:project"`
	TicketHistory []TicketHistory
	gorm.Model
}

// Ticket 表名
func (Ticket) TableName() string {
	return "Ticket"
}

// TicketInfo 工单类型
type TicketInfo struct {
	ID         uint                `json:"id"`
	Number     string              `json:"number"`
	Title      string              `json:"title"`
	Level      uint                `json:"level"`
	Type       string              `json:"type"`
	Stage      string              `json:"stage"`
	Source     string              `json:"source"`
	Improve    string              `json:"improve"`
	Content    string              `json:"content"`
	Solution   string              `json:"solution"`
	StartTime  string              `json:"start_time"`
	EndTime    string              `json:"end_time"`
	Dealer     string              `json:"dealer"`
	PreDealer  string              `json:"pre_dealer"`
	Author     string              `json:"author"`
	Project    string              `json:"project"`
	CreateTime string              `json:"create_time"`
	Histiory   []TicketHistoryInfo `json:"history"`
}

type TicketHistory struct {
	Action   string `gorm:"type:varchar(256);not null;column:action"`
	Content  string `gorm:"type:mediumtext;not null;column:content"`
	From     string `gorm:"type:varchar(256);not null;column:from"`
	To       string `gorm:"type:varchar(256);not null;column:to"`
	TicketID uint
	gorm.Model
}

// TicketHistory 表名
func (TicketHistory) TableName() string {
	return "TicketHistory"
}

// TicketHistoryInfo 工单类型
type TicketHistoryInfo struct {
	ID         uint   `json:"id"`
	Action     string `json:"action"`
	Content    string `json:"content"`
	From       string `json:"from"`
	To         string `json:"to"`
	CreateTime string `json:"create_time"`
}

// TicketType 工单类型结构体
type TicketType struct {
	Name string `gorm:"type:varchar(32);unique;not null;column:name"`
	gorm.Model
}

// TicketType 表名
func (TicketType) TableName() string {
	return "TicketType"
}

// TicketTypeInfo 工单类型
type TicketTypeInfo struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	CreateTime string `json:"create_time"`
}

// TicketLevel 工单等级结构体
type TicketLevel struct {
	Name  string `gorm:"type:varchar(32);unique;not null;column:name"`
	Level uint   `gorm:"type:int(10);unique;not null;column:level"`
	Time  uint   `gorm:"type:int(10);not null;column:time"`
	gorm.Model
}

// TicketLevel 表名
func (TicketLevel) TableName() string {
	return "TicketLevel"
}

// TicketLevelInfo 工单类型
type TicketLevelInfo struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Level      uint   `json:"level"`
	Time       uint   `json:"time"`
	CreateTime string `json:"create_time"`
}

// TicketSource 工单来源结构体
type TicketSource struct {
	Name string `gorm:"type:varchar(32);unique;not null;column:name"`
	gorm.Model
}

// TicketSource 表名
func (TicketSource) TableName() string {
	return "TicketSource"
}

// TicketSource 工单类型
type TicketSourceInfo struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	CreateTime string `json:"create_time"`
}
