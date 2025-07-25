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
	// TODO: Initialize core service and adapters
	// core := core.NewCarbonIntensityService(repo, provider)

	srv := &http.Server{Addr: ":8080"}

	// TODO: Initialize HTTP handler
	// h := handler_http.NewHandler(core)
	// http.Handle("/", h)

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		log.Print("The service is shutting down...")
		srv.Shutdown(context.Background())
	}()

	log.Print("Carbon Intensity Provider service listening...")
	srv.ListenAndServe()
	log.Print("Done")
}