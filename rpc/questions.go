package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetQuestions(c *gin.Context) {
	db := Conn()
	defer db.Close()

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "200"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	// 读取
	var questions []Question
	var Total int
	//db.Where(1)
	ExamineId := c.Query("examine_id")
	if ExamineId != "" {
		db = db.Where("examine_id = ?", ExamineId)
	}

	db.Model(&Question{}).Count(&Total).Limit(limit).Offset(offset).Find(&questions)

	c.JSON(http.StatusOK, Response{RespCode: RespStatusOK, RespDesc: "Success",
		RespData: Item{Limit: uint(limit), Offset: uint(offset), Count: uint(len(questions)), Total: uint(Total), Data: questions}})
}

func CreateQuestion(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "50")

	c.String(http.StatusOK, "Hello %s %s", page, pageSize)
}

func UpdateQuestion(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "50")

	c.String(http.StatusOK, "Hello %s %s", page, pageSize)
}

func DeleteQuestion(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "50")

	c.String(http.StatusOK, "Hello %s %s", page, pageSize)
}

func SearchQuestions(c *gin.Context) {
	db := Conn()
	defer db.Close()

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "200"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	keyword := c.PostForm("keyword")

	// 读取
	var questions []Question
	var Total int

	log.Println(fmt.Sprintf("%%%s%%", keyword))
	if keyword != "" {
		db = db.Where("question_title LIKE ?", fmt.Sprintf("%%%s%%", keyword))
	}

	db.Model(&Question{}).Count(&Total).Limit(limit).Offset(offset).Find(&questions)

	c.JSON(http.StatusOK, Response{RespCode: RespStatusOK, RespDesc: "Success",
		RespData: Item{Limit: uint(limit), Offset: uint(offset), Count: uint(len(questions)), Total: uint(Total), Data: questions}})
}
