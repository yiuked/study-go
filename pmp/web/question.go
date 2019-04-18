package web

import (
	"awesomeProject/pmp/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func Question(w http.ResponseWriter, r *http.Request)  {
	row := db.GetOne("SELECT * FROM `pm_cto_question`")
	fmt.Println(row)
	fmt.Println(RemoteIp(r))
	fmt.Println("打印Header参数列表：")
	if len(r.Header) > 0 {
		for k,v := range r.Header {
			fmt.Printf("%s=%s\n", k, v[0])
		}
	}
	fmt.Println("打印Form参数列表：")
	r.ParseForm()
	if len(r.Form) > 0 {
		for k,v := range r.Form {
			fmt.Printf("%s=%s\n", k, v[0])
		}
	}
	//验证用户名密码，如果成功则header里返回session，失败则返回StatusUnauthorized状态码
	w.WriteHeader(http.StatusOK)
	if (r.Form.Get("user") == "admin") && (r.Form.Get("pass") == "888") {
		w.Write([]byte("hello,验证成功！"))
	} else {
		w.Write([]byte("hello,验证失败了！"))
	}
	data, _ := json.Marshal(row)
	w.Write([]byte(string(data)))
}