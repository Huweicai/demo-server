package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func main() {
	go http.ListenAndServe(":80", Handler{})
	select {}
}

type Handler struct {
}

const template = `
<html>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	<h1 style="text-align:center;white-space: pre-wrap; color:red;font-size:80px">
        
Hello World
服务端IP: %s
客户端地址：%s
UA: %s
</h1>
</html>
`

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "content-type: text/html; charset=utf-8")
	writer.WriteHeader(200)

	addrText := ""
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				addrText += ipnet.IP.To4().String() + ","
			}
		}
	}
	addrText = strings.TrimSuffix(addrText, ",")

	resp := fmt.Sprintf(template, addrText, request.RemoteAddr, request.Header.Get("User-Agent"))
	writer.Write([]byte(resp))
}
