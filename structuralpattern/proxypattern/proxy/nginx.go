package proxy

import (
	"fmt"
	"time"

	application "com.systemdesignpattern.go.proxypattern/main_application"
)

type Nginx struct {
	application          *application.WebServer
	maxRequestsPerMinute int
	requestCount         map[string]int
}

func NewNginx() *Nginx {
	return &Nginx{
		application:          &application.WebServer{},
		maxRequestsPerMinute: 2,
		requestCount:         make(map[string]int),
	}
}

func (n *Nginx) HandleRequest(request string) string {
	if !n.checkRateLimit(request) {
		return "Rate limit exceeded for request: " + request
	}
	return n.application.HandleRequest(request)
}

func (n *Nginx) checkRateLimit(request string) bool {
	key := fmt.Sprintf("%s:%s", request, time.Now().Truncate(time.Minute))
	val, ok := n.requestCount[key]
	if !ok {
		n.requestCount[key] = 1
		return true
	}

	if val+1 > n.maxRequestsPerMinute {
		return false
	}
	n.requestCount[key] = val + 1
	return true
}
