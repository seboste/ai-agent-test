package ports

import "services/job/model"

// Repository is an interface for a database that stores and retrieves jobs.
type Repository interface {
	// Store stores a job in the database.
	Store(job *model.Job) error
	// Find retrieves a job from the database by its ID.
	Find(id string) (*model.Job, error)
	// List retrieves all jobs from the database by status.
	List(status string) ([]*model.Job, error)
}
