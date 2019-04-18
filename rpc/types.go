package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func GetTypes(c *gin.Context)  {
	//page := c.DefaultQuery("page", "1")
	//pageSize := c.DefaultQuery("pageSize", "50")

	db := conn()
	defer db.Close()

	// 读取
	var typeObject Type
	db.First(&typeObject)

	log.Print(typeObject)
	c.String(http.StatusOK, "Hello %s", typeObject.title)
}

func CreateType(c *gin.Context)  {
	db := conn()
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Type{})

	// 创建
	db.Create(&Type{examine_id: 10001, title: "测试1号"})
}

func UpdateType(c *gin.Context){
	db := conn()
	defer db.Close()

	// 读取
	var typeObject Type

	// 更新 - 更新product的price为2000
	db.Model(&typeObject).Update("title", "测试1号 + 1")
}

func DeleteType(c *gin.Context)  {
	db := conn()
	defer db.Close()

	// 读取
	var typeObject Type

	// 删除 - 删除product
	db.Delete(&typeObject)
}

func conn() *gorm.DB{
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "pm_" + defaultTableName;
	}

	db, err := gorm.Open("mysql", "root:@tcp(localhost:3308)/pmp?charset=utf8")
	if err != nil {
		panic("failed to connect database")
	}
	return db;
}