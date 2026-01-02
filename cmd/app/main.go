package main

import (
	"log"
	"net/http"

	"web-based/internal/config"
	"web-based/internal/middleware"
	"web-based/internal/router"
)

func main() {
	cfg := config.Load()

	r := router.New()

	handlerChain := middleware.Recover(
		middleware.Logger(r),
	)

	log.Printf("%s running on :%s\n", cfg.AppName, cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handlerChain))
}
