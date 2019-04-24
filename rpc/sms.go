package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"reflect"
)

type SmsAction struct {
}

func Send(c *gin.Context) {
	smsPhone := c.Query("phone")
	smsType := c.Query("type")

	action := &SmsAction{}
	value := reflect.ValueOf(action)

	arg := reflect.ValueOf(smsPhone)
	method := value.MethodByName(smsType)
	if method.IsValid() {
		result := method.Call([]reflect.Value{arg})
		c.JSON(http.StatusOK, Response{RespCode: RespStatusOK, RespDesc: "Success", RespData: result})
	} else {
		c.JSON(http.StatusOK, Response{RespCode: RespStatusArgs, RespDesc: "Success", RespData: nil})
	}

}

func (s *SmsAction) SendReg(smsPhone string) *gorm.DB{
	db := Conn()
	defer db.Close()

	var lastSms Sms
	db.Where("phone=? AND type=?", smsPhone, SmsTypeReg).First(&lastSms)

	sms := Sms{Type: SmsTypeReg, Phone: smsPhone, Status: SmsStatusInit, Code: "3052", Msg: "尊敬的用户，您的的短信验证为3052."}

	return db.Create(&sms)
}
