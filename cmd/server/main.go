package main

import (
	"log"
	"net/http"

	"github.com/henryarin/portfolio-backend-go/internal/config"
	"github.com/henryarin/portfolio-backend-go/internal/middleware"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	handler := middleware.CORS(cfg.AllowedOrigin, mux)

	log.Println("listening on :" + cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler))
}
