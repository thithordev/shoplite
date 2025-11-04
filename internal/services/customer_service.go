package services

import (
	"shoplite/internal/models"
	"shoplite/internal/repositories"

	"github.com/go-playground/validator/v10"
)

type CustomerService interface {
	Create(input CreateCustomerInput) (*models.Customer, error)
	List() ([]models.Customer, error)
	Get(id uint) (*models.Customer, error)
}

type customerService struct {
	repo      repositories.CustomerRepository
	validator *validator.Validate
}

type CreateCustomerInput struct {
	Name  string `json:"name" validate:"required,min=2,max=120"`
	Email string `json:"email" validate:"required,email"`
}

func NewCustomerService(repo repositories.CustomerRepository, v *validator.Validate) CustomerService {
	return &customerService{repo: repo, validator: v}
}

func (s *customerService) Create(input CreateCustomerInput) (*models.Customer, error) {
	if err := s.validator.Struct(input); err != nil {
		return nil, err
	}
	c := &models.Customer{Name: input.Name, Email: input.Email}
	if err := s.repo.Create(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *customerService) List() ([]models.Customer, error) {
	return s.repo.FindAll()
}

func (s *customerService) Get(id uint) (*models.Customer, error) {
	return s.repo.FindByID(id)
}
