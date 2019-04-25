package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type SmsAction struct {
}

type RequestData struct {
	to         string
	appId      string
	templateId string
	datas      []string
}

func Send(c *gin.Context) {
	smsPhone := c.Query("phone")
	smsType := c.Query("type")

	if IsPhone(smsPhone) {
		c.JSON(http.StatusOK, Response{RespCode: RespStatusArgs, RespDesc: "手机格式错误", RespData: nil})
	}

	action := &SmsAction{}
	value := reflect.ValueOf(action)

	arg := reflect.ValueOf(smsPhone)
	method := value.MethodByName(smsType)
	if method.IsValid() {
		method.Call([]reflect.Value{arg})
		c.JSON(http.StatusOK, Response{RespCode: RespStatusOK, RespDesc: "Success", RespData: nil})
	} else {
		c.JSON(http.StatusOK, Response{RespCode: RespStatusArgs, RespDesc: "参数错误", RespData: nil})
	}

}

func (s *SmsAction) SendReg(smsPhone string) {
	db := Conn()
	defer db.Close()

	var lastSms Sms
	db.Where("phone=? AND type=?", smsPhone, SmsTypeReg).First(&lastSms)
	Code := Rand(100000, 999999)
	sms := Sms{Type: SmsTypeReg, Phone: smsPhone, Status: SmsStatusInit, Code: string(Code), Msg: fmt.Sprintf("尊敬的用户，您的的短信验证为%d.", Code)}
	db.Create(&sms)

	s.sendSms(RequestData{to:smsPhone,templateId:"49324",datas:[]string{string(Code)}})
}

func (s *SmsAction) sendSms(req RequestData) {
	sid, _ := Config.Get("sid")
	url, _ := Config.Get("api")
	appId, _ := Config.Get("api")
	token, _ := Config.Get("token")

	timer := time.Now().Format("20060102150405")
	sig := strings.ToUpper(Md5(fmt.Sprintf("%s%s%s", appId, token, timer)))

	url = fmt.Sprintf("%s/2013-12-26/Accounts/%s/SMS/TemplateSMS?sig=%s", url, sid, sig)

	authorization := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", sid, timer)))
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json;charset=utf-8", "Authorization": authorization}

	req.appId = appId
	jsonReq, _ := json.Marshal(req)

	result := HttpPost(url, string(jsonReq), headers)
	log.Println(result)
}
