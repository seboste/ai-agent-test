package main

import (
	"log"
	"net/http"

	handler_http "github.com/seboste/ai-agent-test/services/job/adapters/handler-http"
	"github.com/seboste/ai-agent-test/services/job/core"
	"github.com/seboste/ai-agent-test/services/job/ports"
)

func main() {
	var service ports.API
	service = core.NewService()
	handler := handler_http.NewHandler(service)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
