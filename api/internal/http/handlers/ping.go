package handlers

import (
	"net/http"
)

func PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("pong")
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("pong"))
	}
}
