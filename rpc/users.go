package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	db := Conn()
	phone := c.Query("phone")
	code := c.Query("code")
	password := c.Query("password")

	if len(phone) < 10 || len(code) < 6 || len(password) < 6 {
		c.JSON(http.StatusOK, Response{RespCode: RespStatusArgs, RespDesc: "参数错误", RespData: nil})
		return
	}

	if !IsPhone(phone) {
		c.JSON(http.StatusOK, Response{RespCode: RespStatusArgs, RespDesc: "请输入正确的手机号码", RespData: nil})
		return
	}

	var exists int
	db.Model(&User{}).Where("phone=?", phone).Count(&exists)
	if exists > 0 {
		c.JSON(http.StatusOK, Response{RespCode: RespStatusArgs, RespDesc: "该手机号码已被注册", RespData: nil})
		return
	}

	var lastSms Sms
	db.Where("phone=? AND type=?", phone, SmsTypeReg).Order("id DESC").First(&lastSms)
	if lastSms.Code != code && lastSms.Status == SmsStatusInit {
		c.JSON(http.StatusOK, Response{RespCode: RespStatusArgs, RespDesc: "验证码错误", RespData: nil})
		return
	}

	user := User{Phone: phone, Status: UserStaAvailable, Level: 0, Password: Md5(password), NickName: "学员"}
	rs := db.Create(&user)

	if rs.RowsAffected > 0 {
		lastSms.Status = SmsStatusVerified
		db.Save(&lastSms)
	}
	c.JSON(http.StatusOK, Response{RespCode: RespStatusOK, RespDesc: "success", RespData: nil})
}

func Login(c *gin.Context) {
	db := Conn()
	phone := c.Query("phone")
	password := c.Query("password")

	if len(phone) < 10 || len(password) < 6 {
		c.JSON(http.StatusOK, Response{RespCode: RespStatusArgs, RespDesc: "参数错误", RespData: nil})
		return
	}

	var user User
	db.Where("phone=?", phone).First(&user)
	if user.Password != Md5(password) {
		c.JSON(http.StatusOK, Response{RespCode: RespStatusArgs, RespDesc: "用户不存在或验证错误", RespData: nil})
		return
	}

	user.LoginCnt += 1
	db.Save(&user)

	var token Token
	db.Where("user_id=?", user.UserId).First(&token)
	// 生成新的过期时间
	tokenExp, _ := Config.Get("user.token_exp")
	tokenExpDuration, _ := time.ParseDuration(tokenExp)
	token.ExpireAt = time.Now().Add(tokenExpDuration)
	// 生成新的token
	token.Token = Md5(fmt.Sprintf("%d%s", user.UserId, time.Now().String()))
	// 创建或者更新用户token
	if token.UserId > 0 {
		db.Save(&token)
	} else {
		token.UserId = user.UserId
		db.Create(&token)
	}

	c.JSON(http.StatusOK, Response{RespCode: RespStatusOK, RespDesc: "success", RespData: token})
}

func IsLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := Conn()
		tokenStr := c.Request.URL.Query().Get("token")
		var token Token
		db.Where("token=?", tokenStr).First(&token)
		if token.ExpireAt.Before(time.Now()) {
			c.JSON(http.StatusForbidden, "Invalid API token")
			c.Abort()
			return
		} else {
			c.
			c.Next()
		}
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := Conn()
		tokenStr := c.Query("token")
		var token Token
		db.Where("token=?", tokenStr).First(&token)
		if token.ExpireAt.Before(time.Now()) {
			c.JSON(http.StatusForbidden, "Invalid API token")
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
