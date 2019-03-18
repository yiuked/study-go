package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func InitHttp() {
	http.HandleFunc("/", indexHandler)
	error := http.ListenAndServe("127.0.0.1:9637", nil)
	if error != nil {
		log.Fatalf(error.Error())
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(remoteIp(r))
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

}

func remoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("Remote_addr"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}
