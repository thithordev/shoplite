package models

import "time"

type Product struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(160);not null" json:"name"`
	Price     float64   `gorm:"not null" json:"price"`
	Stock     int       `gorm:"not null" json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
