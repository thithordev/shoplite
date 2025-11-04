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

type mockCustomerRepo struct{ mock.Mock }

func (m *mockCustomerRepo) Create(c *models.Customer) error { return m.Called(c).Error(0) }
func (m *mockCustomerRepo) FindAll() ([]models.Customer, error) {
	args := m.Called()
	if v := args.Get(0); v != nil {
		return v.([]models.Customer), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *mockCustomerRepo) FindByID(id uint) (*models.Customer, error) {
	args := m.Called(id)
	if v := args.Get(0); v != nil {
		return v.(*models.Customer), args.Error(1)
	}
	return nil, args.Error(1)
}

var _ repositories.CustomerRepository = (*mockCustomerRepo)(nil)

func TestCustomerService_Create_Success(t *testing.T) {
	repo := new(mockCustomerRepo)
	v := validator.New()
	svc := NewCustomerService(repo, v)

	repo.On("Create", mock.AnythingOfType("*models.Customer")).Return(nil).Once()

	cust, err := svc.Create(CreateCustomerInput{Name: "Jane", Email: "jane@example.com"})
	require.NoError(t, err)
	require.NotNil(t, cust)
	assert.Equal(t, "Jane", cust.Name)
	repo.AssertExpectations(t)
}

func TestCustomerService_Create_ValidationError(t *testing.T) {
	repo := new(mockCustomerRepo)
	v := validator.New()
	svc := NewCustomerService(repo, v)
	_, err := svc.Create(CreateCustomerInput{Name: "", Email: "bad"})
	require.Error(t, err)
}

func TestCustomerService_Create_RepoError(t *testing.T) {
	repo := new(mockCustomerRepo)
	v := validator.New()
	svc := NewCustomerService(repo, v)
	repo.On("Create", mock.AnythingOfType("*models.Customer")).Return(errors.New("db error")).Once()
	_, err := svc.Create(CreateCustomerInput{Name: "Jane", Email: "jane@example.com"})
	require.Error(t, err)
	repo.AssertExpectations(t)
}
