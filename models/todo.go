package models

import "time"

const (
	TodoPriorityVeryHigh = "very-high"
	TodoPriorityHigh     = "high"
	TodoPriorityMedium   = "medium"
	TodoPriorityLow      = "low"
	TodoPriorityVeryLow  = "very-low"
)

type Todo struct {
	ID              uint64    `json:"id"`
	ActivityGroupID uint64    `json:"activity_group_id"`
	Title           string    `json:"title"`
	Priority        string    `json:"priority"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
