package main

import (
	"fmt"
	"time"

	"com.systemdesignpattern.go.proxypattern/proxy"
)

func main() {
	nginx := proxy.NewNginx()
	go test(nginx, "request1")
	go test(nginx, "request2")
	time.Sleep(10 * time.Second)
}

func test(nginx *proxy.Nginx, request string) {
	for i := 0; i < 5; i++ {
		fmt.Println(nginx.HandleRequest(request))
		time.Sleep(1 * time.Second)
	}
}
