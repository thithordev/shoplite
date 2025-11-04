package models

import "time"

type Customer struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(120);not null" json:"name"`
	Email     string    `gorm:"type:varchar(180);uniqueIndex;not null" json:"email"`
	Orders    []Order   `json:"orders,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
