package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := Conn()
		tokenStr := c.Query("token")

		var token Token
		if err := db.Where("token=?", tokenStr).First(&token).Error; err != nil {
			c.JSON(http.StatusForbidden, "Invalid API token,token not found")
			c.Abort()
			return
		}

		if token.ExpireAt.Before(time.Now()) {
			c.JSON(http.StatusForbidden, "Invalid API token,token is expire")
			c.Abort()
			return
		}

		var request Request
		if err := db.Where("request_type=? AND router=?", c.Request.Method, c.Request.URL.Path).First(&request).Error; err != nil {
			c.JSON(http.StatusForbidden, "Invalid API token,request not found")
			c.Abort()
			return
		}

		var userAdmin UserAdmin
		if err := db.Where("user_id=?", token.UserId).First(&userAdmin).Error; err != nil {
			c.JSON(http.StatusForbidden, "Invalid API token,404")
			c.Abort()
			return
		}

		var access Access
		if err := db.Where("rule_id=? and request_id=?", userAdmin.RuleId, request.ID).First(&access).Error; err != nil {
			c.JSON(http.StatusForbidden, "Invalid API token,403")
			c.Abort()
			return
		}

		c.Next()
	}
}
