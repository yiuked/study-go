package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Testing(c *gin.Context)  {
	questionIds := c.PostFormMap("question_id")
	log.Println(questionIds)
}