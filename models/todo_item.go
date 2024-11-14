package models

import "time"

// Note: App will be very basic - otherwise DTOs should be considered to prevent leaky APIs
type TodoItem struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	Priority    int        `json:"priority"`
	UserId      *string    `json:"user_id,omitempty"`
}
