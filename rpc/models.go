package main

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 考题类型
type Type struct {
	gorm.Model
	Title string
}
// 考题
type Question struct {
	gorm.Model
	TypeId uint
	QuestionTitle string
	Analyze string
	Answer string
	ItemA string
	ItemB string
	ItemC string
	ItemD string
}

// 用户
type User struct {
	UserId uint `gorm:"primary_key"`
	NickName string
	Phone string
	Password string
	Status string
	Level int8
	LoginCnt uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 用户Token
type Token struct {
	gorm.Model
	UserId uint `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Token string
	ExpireAt time.Time
}

// 短信
type Sms struct {
	gorm.Model
	Type string
	Phone string
	Status string
	Code string
	Msg string
}

// 用户
type Result struct {
	gorm.Model
	TypeId uint
	UserId uint
	Score uint
	TotalScore uint
}
// 用户
type ResultItem struct {
	ID uint
	UserId uint
	ResultId uint
	QuestionId uint
	UserAnswer string
	IsRight uint
}

// 返回信息格式
type Response struct {
	RespCode uint
	RespDesc string
	RespData interface{}
}

// 还回数组格式
type Item struct {
	Limit uint
	Offset uint
	Data interface{}
	Total uint
	Count uint
}

type G struct {
	User User
	Token Token
}

type TestResult struct {
	Score uint
	TotalScore uint
	Answer map[string]string
}