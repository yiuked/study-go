package main

import "github.com/jinzhu/gorm"

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
	gorm.Model
	UserId int
	NickName string
	Phone string
	Password string
	Identity string
}

// 用户
type Sms struct {
	gorm.Model
	Phone string
	Status bool
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