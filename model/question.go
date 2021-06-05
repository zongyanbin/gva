// 自动生成模板Question
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Question struct {
	global.GVA_MODEL
	Branch_office_id int    `json:"branch_office_id" form:"branch_office_id" gorm:"column:branch_office_id;comment:分支机构;type:int;size:11;"`
	Question_id      int    `json:"question_id" form:"question_id" gorm:"column:question_id;comment:问题编号;type:int;size:11;"`
	Question_name    string `json:"question_name" form:"question_name" gorm:"column:question_name;comment:问题标题;type:varchar(191);size:191;"`
	Direction        string `json:"direction" form:"direction" gorm:"column:direction;comment:说明指导;type:varchar(191);size:191;"`
	Answer_state     int    `json:"answer" form:"answer" gorm:"column:answer;comment:是否必答"`
	Topic_id         string `json:"topic_id" form:"topic_id" gorm:"column:topic_id;comment:题型编号;type:varchar(191);size:191;"`
	Author           string `json:"author" form:"author" gorm:"column:author;comment:发题人;type:varchar(191);size:191;"`
	// 关联模型：一对多
	Question_options []Question_options `gorm:"foreignKey:Question_id;references:ID"`
}
// 自定义结构体测试查询数据
type Question_res struct {
      Question_name  string `json:"question_name" form:"question_name" gorm:"column:question_name;comment:问题标题;type:varchar(191);size:191;"`
      Direction  string `json:"direction" form:"direction" gorm:"column:direction;comment:说明指导;type:varchar(191);size:191;"`
}
func (Question) TableName() string {
	return "question"
}
