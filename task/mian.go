package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 启动worker进程
	Dispatch()
	// 启用WEB进程
	InitHttp()
}
