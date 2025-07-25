package handler_http

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/seboste/ai-agent-test/services/job/ports"
)

type Handler struct {
	service ports.API
	rtr     *mux.Router
}

var _ http.Handler = (*Handler)(nil)

func NewHandler(service ports.API) *Handler {
	h := Handler{service: service, rtr: mux.NewRouter()}
	h.rtr.HandleFunc("/jobs", h.handleGetJobs).Methods("GET")
	h.rtr.HandleFunc("/jobs", h.handleCreateJob).Methods("POST")
	h.rtr.HandleFunc("/jobs/{id}", h.handleGetJob).Methods("GET")
	h.rtr.HandleFunc("/jobs/{id}/outcome", h.handleGetJobOutcome).Methods("GET")
	h.rtr.HandleFunc("/jobs/{id}/update-scheduler", h.handleUpdateJobScheduler).Methods("PATCH")
	h.rtr.HandleFunc("/jobs/{id}/update-workerdaemon", h.handleUpdateJobWorkerDaemon).Methods("PATCH")
	return &h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.rtr.ServeHTTP(w, r) //delegate
}

func (h *Handler) handleGetJobs(w http.ResponseWriter, r *http.Request) {
	status := []string{}
	if statusParam := r.URL.Query().Get("status"); statusParam != "" {
		status = strings.Split(statusParam, ",")
	}
	jobs, err := h.service.GetJobs(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(jobs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}

func (h *Handler) handleCreateJob(w http.ResponseWriter, r *http.Request) {
	var job ports.JobCreate
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdJob, err := h.service.CreateJob(job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdJob)
}

func (h *Handler) handleGetJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "invalid uuid", http.StatusBadRequest)
		return
	}
	job, err := h.service.GetJob(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}

func (h *Handler) handleGetJobOutcome(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "invalid uuid", http.StatusBadRequest)
		return
	}
	outcome, err := h.service.GetJobOutcome(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(outcome)
}

func (h *Handler) handleUpdateJobScheduler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "invalid uuid", http.StatusBadRequest)
		return
	}
	var update ports.JobSchedulerUpdate
	err = json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedJob, err := h.service.UpdateJobScheduler(id, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedJob)
}

func (h *Handler) handleUpdateJobWorkerDaemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "invalid uuid", http.StatusBadRequest)
		return
	}
	var update ports.JobWorkerDaemonUpdate
	err = json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedJob, err := h.service.UpdateJobWorkerDaemon(id, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedJob)
}
