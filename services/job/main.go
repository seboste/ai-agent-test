package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
 
	handler_http "github.com/seboste/ai-agent-test/services/job/adapters/handler-http"
	"github.com/seboste/ai-agent-test/services/job/core"
	"github.com/seboste/ai-agent-test/services/job/ports"
)

func main() {
	// Get the port from environment variable, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create HTTP server
	srv := &http.Server{Addr: ":" + port}
  
  var service ports.API
	service = core.NewService()
	handler := handler_http.NewHandler(service)

	// Simple health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
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
