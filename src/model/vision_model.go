package model

import "time"

type Vision struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Body      string    `gorm:"not null" json:"body"`
	Active    bool      `gorm:"not null" json:"active"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}
