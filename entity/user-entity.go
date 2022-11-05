package entity

import (
	"time"
)

// User represents user table in DB
type User struct {
	ID           uint      `gorm:"primaryKey;" json:"id"`
	Name         string    `json:"name"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    string    `json:"created_by"`
	ModifiedAt   time.Time `json:"modified_at"`
	ModifiedBy   string    `json:"modified_by"`
}
