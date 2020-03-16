package web

type Request interface {
	GetParam(key string) string
}

type Response interface {
	Send(data interface{}) error
	JSON(data interface{}) error
	Render(view string, data interface{}) error
	NotFound()
}
