package main

import (
	"fmt"
)

// 定义接口
type DBinter interface {
	connect(dns string) (result DBinter, err error)
}

// 实现接口
type Mysql struct {
}

func (db *Mysql) connect(dns string) (result DBinter, err error) {
	fmt.Println("Connect MySQL success!")
	return db, nil
}

// 实现接口
type Mssql struct {
}

func (db *Mssql) connect(dns string) (result DBinter, err error) {
	fmt.Println("Connect MSSQL success!")
	return db, nil
}
// 定义应用层操作类型
type DB struct {
	driver map[string]DBinter
}
// 注册
func (db *DB) register(name string, driver DBinter) {
	if db.driver == nil {
		db.driver = make(map[string]DBinter)
	}
	if _, err := db.driver[name]; !err {
		db.driver[name] = driver
	}
}
// 获取
func (db *DB) instance(name string) (driver DBinter, err error) {
	if driver, err := db.driver[name]; err {
		return driver,nil
	}
	return nil,err
}

func main() {
	var db DB
	var mysql Mysql
	var mssql Mssql

	db.register("mysql", &mysql)
	db.register("mssql", &mssql)

	mysqlDb,_ := db.instance("mysql")
	mysqlDb.connect("127.0.0.1")
	mssqlDb,_ := db.instance("mysql")
	mssqlDb.connect("127.0.0.1")
}
