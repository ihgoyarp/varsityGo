package router

import (
	"net/http"

	"web-based/internal/handler"
)

func New() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/health", handler.HealthHandler)
	mux.HandleFunc("/echo", handler.EchoHandler)

	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux
}
