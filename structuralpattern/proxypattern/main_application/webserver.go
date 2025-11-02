package application

import "fmt"

type WebServer struct {
}

func (w *WebServer) HandleRequest(request string) string {
	return fmt.Sprintf("Handling request: %s with 200 status code", request)
}
