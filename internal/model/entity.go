package model

import "time"

type Project struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	Link        string    `json:"link"`
	CreatedAt   time.Time `json:"created_at"`
}
