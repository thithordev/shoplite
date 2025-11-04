package services

import (
	"errors"
	"testing"

	"shoplite/internal/models"
	"shoplite/internal/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockProductRepo struct{ mock.Mock }

func (m *mockProductRepo) Create(p *models.Product) error { return m.Called(p).Error(0) }
func (m *mockProductRepo) FindAll() ([]models.Product, error) {
	args := m.Called()
	if v := args.Get(0); v != nil {
		return v.([]models.Product), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *mockProductRepo) FindByID(id uint) (*models.Product, error) {
	args := m.Called(id)
	if v := args.Get(0); v != nil {
		return v.(*models.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

var _ repositories.ProductRepository = (*mockProductRepo)(nil)

func TestProductService_Create_Success(t *testing.T) {
	repo := new(mockProductRepo)
	v := validator.New()
	svc := NewProductService(repo, v)

	repo.On("Create", mock.AnythingOfType("*models.Product")).Return(nil).Once()

	prod, err := svc.Create(CreateProductInput{Name: "Widget", Price: 2.5, Stock: 3})
	require.NoError(t, err)
	require.NotNil(t, prod)
	assert.Equal(t, 2.5, prod.Price)
	repo.AssertExpectations(t)
}

func TestProductService_Create_ValidationError(t *testing.T) {
	repo := new(mockProductRepo)
	v := validator.New()
	svc := NewProductService(repo, v)
	_, err := svc.Create(CreateProductInput{Name: "W", Price: -1, Stock: -1})
	require.Error(t, err)
}

func TestProductService_Create_RepoError(t *testing.T) {
	repo := new(mockProductRepo)
	v := validator.New()
	svc := NewProductService(repo, v)
	repo.On("Create", mock.AnythingOfType("*models.Product")).Return(errors.New("db error")).Once()
	_, err := svc.Create(CreateProductInput{Name: "Widget", Price: 2.5, Stock: 3})
	require.Error(t, err)
	repo.AssertExpectations(t)
}
