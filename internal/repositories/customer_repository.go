package repositories

import (
	"shoplite/internal/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(c *models.Customer) error
	FindAll() ([]models.Customer, error)
	FindByID(id uint) (*models.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) Create(c *models.Customer) error {
	return r.db.Create(c).Error
}

func (r *customerRepository) FindAll() ([]models.Customer, error) {
	var customers []models.Customer
	if err := r.db.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *customerRepository) FindByID(id uint) (*models.Customer, error) {
	var customer models.Customer
	if err := r.db.First(&customer, id).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}
