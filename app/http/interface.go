package http

type Http struct{}

type iHttp interface {
	Launch()
}

func New(http *Http) iHttp {
	return http
}
