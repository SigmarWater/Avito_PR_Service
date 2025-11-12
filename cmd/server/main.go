package main

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	host              = "0.0.0.0"
	port              = "8080"
	readHeaderTimeout = 5 * time.Second
	shutdowmTimeout   = 10 * time.Second
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(10 * time.Second))
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"status":"ok","service":"pr_api"}`))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	server := http.Server{
		Addr:              net.JoinHostPort(host, port),
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	go func() {
		log.Printf("Server start on %s\n", net.JoinHostPort(host, port))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Error start server: %s\n", err.Error())
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Started graceful shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), shutdowmTimeout)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Printf("Error shutdown: %s\n", err.Error())
	}

	log.Println("Success finished server")
}
