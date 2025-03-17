package models

import (
	"time"
)

type Listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Category    string    `json:"category"`
	Owner       string    `json:"owner"` // Product owner's username
	CreatedAt   time.Time `json:"created_at"`
}
