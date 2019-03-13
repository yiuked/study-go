package controller

import (
	"github.com/gin-gonic/gin"
	"time"
)

type User struct{
	user_id int
	username string
}

func user (q *gin.Context) {
	user := User{337217, "zsjr_337217"}
	q.JSON(200, gin.H{"username" : user.username})
}

func Log (q *gin.Context) {
	q.JSON(200, gin.H{"date":time.Now()})
}

func Hello (q *gin.Context) {
	hello := "Hello"
	hello += "world!"
	q.JSON(200, gin.H{"date":hello})
}

func init()  {
	println("Hello,Console:user")
}