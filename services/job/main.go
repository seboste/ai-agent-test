package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Get the port from environment variable, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create HTTP server
	srv := &http.Server{Addr: ":" + port}

	// Simple health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Placeholder for jobs endpoints - to be implemented with proper handlers
	http.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[]`))
		case http.MethodPost:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(`{"error": "Not implemented"}`))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Method not allowed"}`))
		}
	})

	// Handle graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		log.Print("The job service is shutting down...")
		srv.Shutdown(context.Background())
	}()

	log.Printf("Job service listening on port %s...", port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Print("Job service stopped")
}
