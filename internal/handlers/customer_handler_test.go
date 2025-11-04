package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"shoplite/internal/models"
	"shoplite/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockCustomerService struct{ mock.Mock }

func (m *mockCustomerService) Create(input services.CreateCustomerInput) (*models.Customer, error) {
	args := m.Called(input)
	if v := args.Get(0); v != nil {
		return v.(*models.Customer), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *mockCustomerService) List() ([]models.Customer, error) {
	args := m.Called()
	if v := args.Get(0); v != nil {
		return v.([]models.Customer), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *mockCustomerService) Get(id uint) (*models.Customer, error) {
	args := m.Called(id)
	if v := args.Get(0); v != nil {
		return v.(*models.Customer), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestCustomerHandler_Create_And_List(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := new(mockCustomerService)
	h := NewCustomerHandler(svc)

	r := gin.New()
	r.POST("/customers", h.Create)
	r.GET("/customers", h.List)

	// Mock create
	svc.On("Create", services.CreateCustomerInput{Name: "Jane", Email: "jane@example.com"}).Return(&models.Customer{ID: 1, Name: "Jane", Email: "jane@example.com"}, nil).Once()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/customers", bytes.NewBufferString(`{"name":"Jane","email":"jane@example.com"}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusCreated, w.Code)

	// Mock list
	svc.On("List").Return([]models.Customer{{ID: 1, Name: "Jane", Email: "jane@example.com"}}, nil).Once()
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest(http.MethodGet, "/customers", nil)
	r.ServeHTTP(w2, req2)
	require.Equal(t, http.StatusOK, w2.Code)
}

func TestCustomerHandler_Create_Invalid(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := new(mockCustomerService)
	h := NewCustomerHandler(svc)
	r := gin.New()
	r.POST("/customers", h.Create)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/customers", bytes.NewBufferString(`{"name":"","email":"bad"}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusBadRequest, w.Code)
}
