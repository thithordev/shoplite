package services

import (
	"errors"
	"testing"
	"time"

	"shoplite/internal/models"
	"shoplite/internal/repositories"
	"shoplite/internal/testutil"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type mockOrderRepo struct{ mock.Mock }

func (m *mockOrderRepo) Create(o *models.Order) error { return m.Called(o).Error(0) }
func (m *mockOrderRepo) FindAll() ([]models.Order, error) {
	args := m.Called()
	if v := args.Get(0); v != nil {
		return v.([]models.Order), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *mockOrderRepo) FindByID(id uint) (*models.Order, error) {
	args := m.Called(id)
	if v := args.Get(0); v != nil {
		return v.(*models.Order), args.Error(1)
	}
	return nil, args.Error(1)
}

var _ repositories.OrderRepository = (*mockOrderRepo)(nil)

type mockProductRepo2 struct{ mock.Mock }

func (m *mockProductRepo2) Create(p *models.Product) error { return m.Called(p).Error(0) }
func (m *mockProductRepo2) FindAll() ([]models.Product, error) {
	args := m.Called()
	if v := args.Get(0); v != nil {
		return v.([]models.Product), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *mockProductRepo2) FindByID(id uint) (*models.Product, error) {
	args := m.Called(id)
	if v := args.Get(0); v != nil {
		return v.(*models.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

var _ repositories.ProductRepository = (*mockProductRepo2)(nil)

type mockCustomerRepo2 struct{ mock.Mock }

func (m *mockCustomerRepo2) Create(c *models.Customer) error { return m.Called(c).Error(0) }
func (m *mockCustomerRepo2) FindAll() ([]models.Customer, error) {
	args := m.Called()
	if v := args.Get(0); v != nil {
		return v.([]models.Customer), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *mockCustomerRepo2) FindByID(id uint) (*models.Customer, error) {
	args := m.Called(id)
	if v := args.Get(0); v != nil {
		return v.(*models.Customer), args.Error(1)
	}
	return nil, args.Error(1)
}

var _ repositories.CustomerRepository = (*mockCustomerRepo2)(nil)

func TestOrderService_Create_Success(t *testing.T) {
	db := testutil.NewTestDB(t)
	testutil.TruncateAll(t, db)
	// seed a product in real DB so the insert constraints will pass
	realProd := &models.Product{Name: "Seed", Price: 1.0, Stock: 100}
	require.NoError(t, db.Create(realProd).Error)
	// customer must exist for FK
	realCust := &models.Customer{Name: "Carl", Email: "carl@example.com"}
	require.NoError(t, db.Create(realCust).Error)

	orderRepo := new(mockOrderRepo) // not used in Create
	prodRepo := new(mockProductRepo2)
	custRepo := new(mockCustomerRepo2)

	// Mock the existence checks
	prodRepo.On("FindByID", realProd.ID).Return(realProd, nil).Maybe()
	custRepo.On("FindByID", realCust.ID).Return(realCust, nil).Maybe()

	svc := NewOrderService(db, orderRepo, prodRepo, custRepo, validator.New())

	inp := CreateOrderInput{
		CustomerID: realCust.ID,
		OrderDate:  time.Now(),
		Status:     "pending",
		Items:      []OrderItemInput{{ProductID: realProd.ID, Quantity: 2, Price: 1.0}},
	}
	order, err := svc.Create(inp)
	require.NoError(t, err)
	require.NotNil(t, order)
	require.Equal(t, 1, len(order.Items))
}

func TestOrderService_Create_ProductMissing(t *testing.T) {
	db := testutil.NewTestDB(t)
	testutil.TruncateAll(t, db)
	orderRepo := new(mockOrderRepo)
	prodRepo := new(mockProductRepo2)
	custRepo := new(mockCustomerRepo2)
	custRepo.On("FindByID", uint(1)).Return(&models.Customer{ID: 1, Name: "X", Email: "x@y.z"}, nil).Once()
	prodRepo.On("FindByID", uint(42)).Return(nil, errors.New("not found")).Once()

	svc := NewOrderService(db, orderRepo, prodRepo, custRepo, validator.New())
	_, err := svc.Create(CreateOrderInput{CustomerID: 1, OrderDate: time.Now(), Items: []OrderItemInput{{ProductID: 42, Quantity: 1, Price: 1}}})
	require.Error(t, err)
}

func TestOrderService_ListAndGet_Empty(t *testing.T) {
	db := testutil.NewTestDB(t)
	repo := new(mockOrderRepo)
	repo.On("FindAll").Return([]models.Order{}, nil).Once()
	repo.On("FindByID", uint(1)).Return(nil, gorm.ErrRecordNotFound).Once()
	svc := &orderService{db: db, repo: repo, prodRepo: new(mockProductRepo2), custRepo: new(mockCustomerRepo2), validator: validator.New()}
	list, err := svc.List()
	require.NoError(t, err)
	require.Len(t, list, 0)
	_, err = svc.Get(1)
	require.Error(t, err)
}
