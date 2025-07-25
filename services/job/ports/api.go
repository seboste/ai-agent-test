package ports

import "github.com/google/uuid"

// API is the primary port for the job service
type API interface {
	// GetJobs retrieves a list of jobs, optionally filtered by status.
	GetJobs(status []string) ([]Job, error)
	// CreateJob creates a new job.
	CreateJob(job JobCreate) (Job, error)
	// GetJob retrieves a job by its ID.
	GetJob(id uuid.UUID) (Job, error)
	// GetJobOutcome retrieves the outcome of a job.
	GetJobOutcome(id uuid.UUID) (JobOutcome, error)
	// UpdateJobScheduler updates a job from the scheduler's perspective.
	UpdateJobScheduler(id uuid.UUID, update JobSchedulerUpdate) (Job, error)
	// UpdateJobWorkerDaemon updates a job from the worker daemon's perspective.
	UpdateJobWorkerDaemon(id uuid.UUID, update JobWorkerDaemonUpdate) (Job, error)
}
