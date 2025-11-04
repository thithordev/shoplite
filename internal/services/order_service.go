package services

import (
	"errors"
	"time"

	"shoplite/internal/models"
	"shoplite/internal/repositories"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type OrderService interface {
	Create(input CreateOrderInput) (*models.Order, error)
	List() ([]models.Order, error)
	Get(id uint) (*models.Order, error)
}

type orderService struct {
	db        *gorm.DB
	repo      repositories.OrderRepository
	prodRepo  repositories.ProductRepository
	custRepo  repositories.CustomerRepository
	validator *validator.Validate
}

type OrderItemInput struct {
	ProductID uint    `json:"product_id" validate:"required,gt=0"`
	Quantity  int     `json:"quantity" validate:"required,gt=0"`
	Price     float64 `json:"price" validate:"required,gt=0"`
}

type CreateOrderInput struct {
	CustomerID uint             `json:"customer_id" validate:"required,gt=0"`
	OrderDate  time.Time        `json:"order_date" validate:"required"`
	Status     string           `json:"status" validate:"omitempty,oneof=pending paid shipped cancelled"`
	Items      []OrderItemInput `json:"items" validate:"required,min=1,dive"`
}

func NewOrderService(db *gorm.DB, repo repositories.OrderRepository, prodRepo repositories.ProductRepository, custRepo repositories.CustomerRepository, v *validator.Validate) OrderService {
	return &orderService{db: db, repo: repo, prodRepo: prodRepo, custRepo: custRepo, validator: v}
}

func (s *orderService) Create(input CreateOrderInput) (*models.Order, error) {
	if err := s.validator.Struct(input); err != nil {
		return nil, err
	}
	// Verify customer exists
	if _, err := s.custRepo.FindByID(input.CustomerID); err != nil {
		return nil, err
	}

	order := &models.Order{
		CustomerID: input.CustomerID,
		OrderDate:  input.OrderDate,
		Status:     firstNonEmpty(input.Status, "pending"),
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// Create order
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		// Create items
		for _, it := range input.Items {
			// Optionally ensure product exists
			if _, err := s.prodRepo.FindByID(it.ProductID); err != nil {
				return err
			}
			item := models.OrderItem{
				OrderID:   order.ID,
				ProductID: it.ProductID,
				Quantity:  it.Quantity,
				Price:     it.Price,
			}
			if err := tx.Create(&item).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Load associations
	if err := s.db.Preload("Customer").Preload("Items").Preload("Items.Product").First(order, order.ID).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (s *orderService) List() ([]models.Order, error) {
	return s.repo.FindAll()
}

func (s *orderService) Get(id uint) (*models.Order, error) {
	if id == 0 {
		return nil, errors.New("invalid id")
	}
	return s.repo.FindByID(id)
}

func firstNonEmpty(v string, def string) string {
	if v != "" {
		return v
	}
	return def
}
