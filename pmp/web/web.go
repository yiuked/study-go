package web

import (
	"log"
	"net"
	"net/http"
)

func InitHttp() {
	http.HandleFunc("/", Question)
	error := http.ListenAndServe("127.0.0.1:9637", nil)
	if error != nil {
		log.Fatalf(error.Error())
	}
}

func RemoteIp(req *http.Request) string {
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