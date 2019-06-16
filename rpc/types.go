package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTypes(c *gin.Context) {
	db := Conn()
	defer db.Close()

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	// 读取
	var types []Type
	var Total int
	db.Find(&types)

	db.Model(&Question{}).Count(&Total).Limit(limit).Offset(offset).Find(&types)

	c.JSON(http.StatusOK, Response{RespCode: RespStatusOK, RespDesc: "Success", RespData: types})
}

func CreateType(c *gin.Context) {
	db := Conn()
	defer db.Close()



	// Migrate the schema
	db.AutoMigrate(&Type{})

	// 创建
	db.Create(&Type{Title: "测试1号"})

	c.JSON(http.StatusOK, Response{RespCode: RespStatusOK, RespDesc: "Success", RespData: nil})
}

func UpdateType(c *gin.Context) {
	db := Conn()
	defer db.Close()

	// 读取
	var typeObject Type

	// 更新 - 更新product的price为2000
	db.Model(&typeObject).Update("title", "测试1号 + 1")
}

func DeleteType(c *gin.Context) {
	db := Conn()
	defer db.Close()

	// 读取
	var typeObject Type

	// 删除 - 删除product
	db.Delete(&typeObject)
}