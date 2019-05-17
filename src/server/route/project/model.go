package project

import (
	"mos/src/server/route/jenkins"
	"mos/src/server/route/ticket"

	"github.com/jinzhu/gorm"
)

// Project 项目结构体
type Project struct {
	Name       string `gorm:"type:varchar(32);unique;not null;column:name"`
	Tickekts   []ticket.Ticket
	JenkinsJob []jenkins.JenkinsJob
	gorm.Model
}

// Project 表名
func (Project) TableName() string {
	return "Project"
}

// ProjectInfo 项目信息
type ProjectInfo struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	CreateTime string `json:"create_time"`
}
