package main

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 考题类型
type Type struct {
	Title string
	ExamineId int
}
// 考题
type Question struct {
	QuestionId int
	ExamineId int
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
	UserId int `gorm:"primary_key"`
	NickName string
	Phone string
	Password string
	Status string
	Level int8
	LoginCnt int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 用户Token
type Token struct {
	gorm.Model
	UserId int
	Token string
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

// 返回信息格式
type Response struct {
	RespCode int
	RespDesc string
	RespData interface{}
}

// 还回数组格式
type Item struct {
	Limit int
	Offset int
	Data interface{}
	Total int
	Count int
}