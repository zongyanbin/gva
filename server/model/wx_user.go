package model

import (
	"gin-vue-admin/global"
	"gorm.io/gorm"
)

type WxUser struct {
	global.GVA_MODEL
	gorm.Model
	OpenId 		string		`json:"openid" form:"opneid" gorm:"comment:用户的openid"`
	NickName 	string  	`json:"nick_name" form:"nick_name" gorm:"comment:用户昵称"`
	AvatarUrl 	string		`json:"avatar_url" form:"avatar_url" gorm:"comment:用户头像"`
	Gender 		string		`json:"gender" form:"gender" gorm:"comment:性别  0-男、1-女、2-保密"`
	Country 	string		`json:"country" form:"country" gorm:"comment:所在国家"`
	Province 	string		`json:"province" form:"provice" gorm:"comment:省份"`
	City		string		`json:"cityd" form:"city" gorm:"commit:城市"`
	Language	string		`json:"language" form:"语言" gorm:"commit:语言"`
	Mobile		string		`json:"mobile" form:"电话号码" gorm:"commit:手机号码"`
}

func (WxUser) TableName() string {
	return "wx_users"
}