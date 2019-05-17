package jenkins

import (
	"mos/src/glo"

	"github.com/jinzhu/gorm"
)

// JenkinsJob jenkins构建历史结构体
type JenkinsJob struct {
	Title       string `gorm:"type:varchar(256);not null;column:title"`
	Module      string `gorm:"type:varchar(256);not null;column:module"`
	Tag         string `gorm:"type:varchar(256);not null;column:tag"`
	BuildNumber string `gorm:"type:varchar(256);not null;column:build_number"`
	Hosts       string `gorm:"type:varchar(256);not null;column:hosts"`
	BuildUser   string `gorm:"type:varchar(256);not null;column:build_user"`
	BuildStatus string `gorm:"type:varchar(256);not null;column:build_status"`
	ConsoleURL  string `gorm:"type:varchar(256);not null;column:console_url"`
	Project     string `gorm:"type:varchar(256);not null;column:project"`
	gorm.Model
}

// TableName JenkinsJob 表名
func (JenkinsJob) TableName() string {
	return "JenkinsJob"
}

// Create 插入jenkins记录
func (j *JenkinsJob) Create(p *JobPost) (res bool, err error) {
	j.Title = p.Title
	j.Module = p.Module
	j.Tag = p.Tag
	j.BuildNumber = p.BuildNumber
	j.Hosts = p.Hosts
	j.BuildUser = p.BuildUser
	j.BuildStatus = p.BuildStatus
	j.ConsoleURL = p.ConsoleURL
	j.Project = p.Project
	// log.Printf("%s, %s, %s", p.Title, p.Module, p.Tag)
	if err = glo.Db.Create(&j).Error; err != nil {
		return false, err
	}
	return true, nil
}

// JenkinsPostForm 上报post参数
type JobPost struct {
	Title       string `json:"title"`
	Project     string `json:"project"`
	Module      string `json:"module"`
	Tag         string `json:"tag"`
	BuildNumber string `json:"build_number"`
	Hosts       string `json:"hosts"`
	BuildUser   string `json:"build_user"`
	BuildStatus string `json:"build_status"`
	ConsoleURL  string `json:"console_url"`
}

// JenkinsJobInfo 项目信息
type JenkinsJobInfo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Project     string `json:"project"`
	Module      string `json:"module"`
	Tag         string `json:"tag"`
	BuildNumber string `json:"build_number"`
	Hosts       string `json:"hosts"`
	BuildUser   string `json:"build_user"`
	BuildStatus string `json:"build_status"`
	ConsoleURL  string `json:"console_url"`
	CreateTime  string `json:"create_time"`
}
