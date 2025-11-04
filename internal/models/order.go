package models

import "time"

type Order struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	CustomerID uint        `json:"customer_id"`
	Customer   Customer    `json:"customer"`
	OrderDate  time.Time   `json:"order_date"`
	Status     string      `gorm:"type:varchar(50);not null;default:'pending'" json:"status"`
	Items      []OrderItem `json:"items"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}
