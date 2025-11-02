package application

type IServer interface {
	HandleRequest(request string) string
}
