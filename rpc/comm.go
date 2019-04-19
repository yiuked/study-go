package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
