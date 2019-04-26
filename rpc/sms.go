package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type SmsAction struct {
}

type ReqData struct {
	To         string   `json:"to"`
	AppId      string   `json:"appId"`
	TemplateId string   `json:"templateId"`
	Datas      []string `json:"datas"`
}

type ResData struct {
	StatusCode  string
	TemplateSMS map[string]string
}

func Send(c *gin.Context) {
	smsPhone := c.Query("phone")
	smsType := c.Query("type")

	if IsPhone(smsPhone) == false {
		c.JSON(http.StatusOK, Response{RespCode: RespStatusArgs, RespDesc: "手机格式错误", RespData: nil})
	}

	action := &SmsAction{}
	value := reflect.ValueOf(action)

	arg := reflect.ValueOf(smsPhone)
	method := value.MethodByName(smsType)
	if method.IsValid() {
		rs := method.Call([]reflect.Value{arg})
		if rs[0].Interface() == true {
			c.JSON(http.StatusOK, Response{RespCode: RespStatusOK, RespDesc: "Success", RespData: nil})
		} else {
			error :=  rs[1].Interface().(error)
			c.JSON(http.StatusOK, Response{RespCode: RespStatusSend, RespDesc: error.Error(), RespData:  nil})
		}
	} else {
		c.JSON(http.StatusOK, Response{RespCode: RespStatusArgs, RespDesc: "参数错误", RespData: nil})
	}
}

func (s *SmsAction) SendReg(smsPhone string) (bool, error) {
	db := Conn()
	defer db.Close()

	var lastSms Sms
	db.Where("phone=? AND type=?", smsPhone, SmsTypeReg).Order("id DESC").First(&lastSms)
	if time.Now().Sub(lastSms.CreatedAt).Seconds() < 60 {
		return false, errors.New("操作过于频繁")
	}

	var count int
	db.Model(&Sms{}).Where("phone=? AND created_at>?", smsPhone,time.Now().Format("2006-01-02")).Count(&count)
	if count >= 10 {
		return false, errors.New("当日操作已达到上限")
	}

	Code := strconv.Itoa(Rand(100000, 999999))
	sms := Sms{Type: SmsTypeReg, Phone: smsPhone, Status: SmsStatusInit, Code: Code, Msg: fmt.Sprintf("尊敬的用户，您的的短信验证为%s.", Code)}
	db.Create(&sms)

	var replaceData = []string{Code}
	reqData := ReqData{To: smsPhone, TemplateId: "49324", Datas: replaceData}

	resData := s.sendSms(reqData)
	if resData.StatusCode != "000000" {
		return false,errors.New("错误码:" + resData.StatusCode)
	}
	return true,nil
}

func (s *SmsAction) sendSms(req ReqData) ResData {
	// 配置conf.yaml中的配置信息
	sid, _ := Config.Get("sms.sid")
	url, _ := Config.Get("sms.api")
	appId, _ := Config.Get("sms.app_id")
	token, _ := Config.Get("sms.token")

	// 获取当前时间，并生成请求签名
	timer := time.Now().Format("20060102150405")
	sig := strings.ToUpper(Md5(fmt.Sprintf("%s%s%s", sid, token, timer)))

	// 生成请求的URL地址
	url = fmt.Sprintf("%s/2013-12-26/Accounts/%s/SMS/TemplateSMS?sig=%s", url, sid, sig)
	log.Println(url)

	// 组织请求头信息
	authorization := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", sid, timer)))
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json;charset=utf-8", "Authorization": authorization}

	req.AppId = appId

	// 将请求数据转为JSON
	jsonReq, errs := json.Marshal(req) //转换成JSON返回的是byte[]
	if errs != nil {
		log.Println(errs.Error())
	}
	log.Println(string(jsonReq)) //byte[]转换成string 输出

	// 发起短信请求
	result := HttpPost(url, string(jsonReq), headers)
	log.Println(result)

	resData := ResData{}
	json.Unmarshal([]byte(result), &resData)
	return resData
}
