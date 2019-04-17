package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetQuestions(c *gin.Context)  {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "50")

	c.String(http.StatusOK, "Hello %s %s", page, pageSize)
}

func CreateQuestion(c *gin.Context)  {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "50")

	c.String(http.StatusOK, "Hello %s %s", page, pageSize)
}

func UpdateQuestion(c *gin.Context){
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "50")

	c.String(http.StatusOK, "Hello %s %s", page, pageSize)
}

func DeleteQuestion(c *gin.Context)  {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "50")

	c.String(http.StatusOK, "Hello %s %s", page, pageSize)
}