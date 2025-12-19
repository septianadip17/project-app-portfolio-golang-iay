package model

import "time"

// Sesuai tabel projects
type Project struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	Link        string    `json:"link"`
	CreatedAt   time.Time `json:"created_at"`
}

// Sesuai tabel experiences
type Experience struct {
	ID          int        `json:"id"`
	Role        string     `json:"role"`
	Company     string     `json:"company"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date"` // Pointer karena bisa NULL (jika masih bekerja)
	Description string     `json:"description"`
}

// Sesuai tabel contacts
type Contact struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Message     string    `json:"message"`
	SubmittedAt time.Time `json:"submitted_at"`
}
