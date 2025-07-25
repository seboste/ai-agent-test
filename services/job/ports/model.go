package ports

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	ID              uuid.UUID              `json:"id"`
	UserID          uuid.UUID              `json:"userId"`
	JobName         string                 `json:"jobName"`
	Image           ContainerImage         `json:"image"`
	Parameters      map[string]string      `json:"parameters"`
	CreationZone    string                 `json:"creationZone"`
	Status          string                 `json:"status"`
	CreatedAt       time.Time              `json:"createdAt"`
	UpdatedAt       time.Time              `json:"updatedAt"`
	Result          string                 `json:"result"`
	WorkerID        *uuid.UUID             `json:"workerId"`
	ErrorMessage    string                 `json:"errorMessage"`
	ComputeZone     string                 `json:"computeZone"`
	CarbonIntensity int                    `json:"carbonIntensity"`
	CarbonSavings   int                    `json:"carbonSavings"`
}

type JobCreate struct {
	JobName      string            `json:"jobName"`
	CreationZone string            `json:"creationZone"`
	Image        ContainerImage    `json:"image"`
	Parameters   map[string]string `json:"parameters"`
}

type ContainerImage struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type JobOutcome struct {
	JobName         string `json:"jobName"`
	Status          string `json:"status"`
	Result          string `json:"result"`
	ErrorMessage    string `json:"errorMessage"`
	ComputeZone     string `json:"computeZone"`
	CarbonIntensity int    `json:"carbonIntensity"`
	CarbonSavings   int    `json:"carbonSavings"`
}

type JobSchedulerUpdate struct {
	WorkerID        uuid.UUID `json:"workerId"`
	ComputeZone     string    `json:"computeZone"`
	CarbonIntensity int       `json:"carbonIntensity"`
	CarbonSavings   int       `json:"carbonSavings"`
	Status          JobStatus `json:"status"`
}

type JobStatus string

const (
	JobStatusQueued    JobStatus = "queued"
	JobStatusScheduled JobStatus = "scheduled"
	JobStatusRunning   JobStatus = "running"
	JobStatusCompleted JobStatus = "completed"
	JobStatusFailed    JobStatus = "failed"
	JobStatusCancelled JobStatus = "cancelled"
)

type JobWorkerDaemonUpdate struct {
	Status       JobStatus `json:"status"`
	Result       string    `json:"result"`
	ErrorMessage string    `json:"errorMessage"`
}
