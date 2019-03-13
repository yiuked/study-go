package main

import "github.com/gin-gonic/gin"
import _ "awesomeProject/controller"

func main() {
	r := gin.Default()
	//r.GET("/log", ctr.Log)
	//r.GET("/hello", ctr.Hello)
	r.Run()
}
