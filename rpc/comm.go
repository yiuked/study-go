package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/rand"
	"regexp"
)

func Conn() *gorm.DB {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "pm_" + defaultTableName;
	}

	db, err := gorm.Open("mysql", "root:@tcp(localhost:3308)/pmp?charset=utf8")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Rand(min int , max int) int {
	return min + rand.Intn(max-min)
}
// 判断是否为手机号
func IsPhone(phone string)  bool{
	reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phone)
}