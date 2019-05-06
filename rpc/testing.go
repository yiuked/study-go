package main

import (
	"github.com/gin-gonic/gin"
)

func Testing(c *gin.Context)  {
	questionIds := c.PostFormMap("question_id")
}