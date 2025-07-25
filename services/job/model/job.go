package model

import "time"

// Job represents a job in the system.
type Job struct {
	ID          string
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
}
