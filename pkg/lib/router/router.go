package router

import "net/http"

type Router interface {
	GET(uri string, f http.Handler)
	POST(uri string, f http.Handler)
	Handler() http.Handler
}

type defaultRouter struct {
	router *http.ServeMux
}

func NewDefaultRouter() Router {
	return &defaultRouter{router: http.NewServeMux()}
}

func (r *defaultRouter) GET(uri string, f http.Handler) {
	r.router.Handle(uri, f)
}

func (r *defaultRouter) POST(uri string, f http.Handler) {
	r.router.Handle(uri, f)
}

func (r *defaultRouter) Handler() http.Handler {
	return r.router
}
