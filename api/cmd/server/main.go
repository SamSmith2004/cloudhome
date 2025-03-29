package main

import (
	"cloudhome/config"
	"cloudhome/internal/http/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	r := router.New()
	r.RegisterRoutes()

	serverAddr := fmt.Sprintf(":%d", cfg.Server.Port)
	server := &http.Server{
		Addr:    serverAddr,
		Handler: r.Handler(),
	}

	log.Printf("Starting server on %s", serverAddr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
