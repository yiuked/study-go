package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetResults(c *gin.Context)  {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "50")

	c.String(http.StatusOK, "Hello %s %s", page, pageSize)
}

func CreateResult(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "50")

	c.String(http.StatusOK, "Hello %s %s", page, pageSize)
}

func UpdateResult(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "50")

	c.String(http.StatusOK, "Hello %s %s", page, pageSize)
}

func DeleteResult(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "50")

	c.String(http.StatusOK, "Hello %s %s", page, pageSize)
}
