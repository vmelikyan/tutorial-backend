package model

import (
	"time"
)

type Task struct {
	ID        int        `json:"id"`
	Task      string     `json:"task"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Deleted   bool       `json:"deleted"`
}
