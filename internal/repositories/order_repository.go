package repositories

import (
	"shoplite/internal/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(o *models.Order) error
	FindAll() ([]models.Order, error)
	FindByID(id uint) (*models.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(o *models.Order) error {
	return r.db.Create(o).Error
}

func (r *orderRepository) FindAll() ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.Preload("Customer").Preload("Items").Preload("Items.Product").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) FindByID(id uint) (*models.Order, error) {
	var order models.Order
	if err := r.db.Preload("Customer").Preload("Items").Preload("Items.Product").First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
