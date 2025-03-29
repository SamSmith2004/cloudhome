package router

import (
	"cloudhome/internal/http/handlers"
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func New() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) Handler() http.Handler {
	return r.mux
}

func (r *Router) RegisterRoutes() {
	r.mux.HandleFunc("/ping", handlers.PingHandler())
	r.mux.HandleFunc("POST /hello", handlers.HelloHandler())
}
