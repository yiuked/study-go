package main

type Type struct {
	Title string
	ExamineId int
}

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

type Response struct {
	RespCode int
	RespDesc string
	RespData interface{}
}

type Item struct {
	Limit int
	Offset int
	Data interface{}
	Total int
	Count int
}