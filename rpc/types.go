package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTypes(c *gin.Context)  {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "50")

	c.String(http.StatusOK, "Hello %s %s", page, pageSize)
}

func CreateType(c *gin.Context)  {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "50")

	c.String(http.StatusOK, "Hello %s %s", page, pageSize)
}

func UpdateType(c *gin.Context){
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "50")

	c.String(http.StatusOK, "Hello %s %s", page, pageSize)
}

func DeleteType(c *gin.Context)  {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "50")

	c.String(http.StatusOK, "Hello %s %s", page, pageSize)
}