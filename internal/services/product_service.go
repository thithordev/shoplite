package services

import (
	"shoplite/internal/models"
	"shoplite/internal/repositories"

	"github.com/go-playground/validator/v10"
)

type ProductService interface {
	Create(input CreateProductInput) (*models.Product, error)
	List() ([]models.Product, error)
	Get(id uint) (*models.Product, error)
}

type productService struct {
	repo      repositories.ProductRepository
	validator *validator.Validate
}

type CreateProductInput struct {
	Name  string  `json:"name" validate:"required,min=2,max=160"`
	Price float64 `json:"price" validate:"required,gt=0"`
	Stock int     `json:"stock" validate:"required,gte=0"`
}

func NewProductService(repo repositories.ProductRepository, v *validator.Validate) ProductService {
	return &productService{repo: repo, validator: v}
}

func (s *productService) Create(input CreateProductInput) (*models.Product, error) {
	if err := s.validator.Struct(input); err != nil {
		return nil, err
	}
	p := &models.Product{Name: input.Name, Price: input.Price, Stock: input.Stock}
	if err := s.repo.Create(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *productService) List() ([]models.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) Get(id uint) (*models.Product, error) {
	return s.repo.FindByID(id)
}
