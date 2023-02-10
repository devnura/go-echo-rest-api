package entity

import (
	"time"
)

// User represents user table in DB
type Question struct {
	ID         uint      `gorm:"primaryKey;" json:"id"`
	Question   string    `gorm:"column:question" json:"question"`
	UserId     string    `gorm:"column:user_id" json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}
