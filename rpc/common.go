package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func Conn() *gorm.DB {
	userName, _ := Config.Get("db.username")
	password, _ := Config.Get("db.password")
	host, _ := Config.Get("db.host")
	port, _ := Config.Get("db.port")
	database, _ := Config.Get("db.database")
	charset, _ := Config.Get("db.charset")
	prefix, _ := Config.Get("db.prefix")

	// 设置数据库表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return prefix + defaultTableName
	}

	// 建立数据库连接
	db, err := gorm.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Asia%%2FShanghai",
		userName, password, host, port, database, charset))

	if err != nil {
		panic(err)
		panic("failed to connect database")
	}

	return db.Debug()
}

func Rand(min int, max int) int {
	return min + rand.Intn(max-min)
}

// 判断是否为手机号
func IsPhone(phone string) bool {
	reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phone)
}

func HttpGet(url string, headers map[string]string) string {
	req, reqErr := http.NewRequest("GET", url, nil)
	if reqErr != nil {
		log.Println(reqErr.Error())
	}

	if len(headers) > 0 {
		for header, value := range headers {
			req.Header.Add(header, value)
		}
	}

	res, doErr := http.DefaultClient.Do(req)
	if doErr != nil {
		log.Println(doErr.Error())
	}

	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Println(readErr.Error())
	}

	return string(body)
}

func HttpPost(url string, data string, headers map[string]string) string {
	payload := strings.NewReader(data)

	req, reqErr := http.NewRequest("POST", url, payload)
	if reqErr != nil {
		log.Println(reqErr.Error())
	}
	if len(headers) > 0 {
		for header, value := range headers {
			req.Header.Add(header, value)
		}
	}

	res, doErr := http.DefaultClient.Do(req)
	if doErr != nil {
		log.Println(doErr.Error())
	}
	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Println(readErr.Error())
	}

	return string(body)
}

func ParseUrl(data map[string]string) string {
	var param []string
	for key, value := range data {
		param = append(param, fmt.Sprintf("%s=%s", key, value))
	}
	return strings.Join(param, "&")
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func StrToTimer(value string) time.Time {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai")
	valueTimer, _ := time.ParseInLocation("2006-01-02 15:04:05", value, cstSh)
	return valueTimer
}
