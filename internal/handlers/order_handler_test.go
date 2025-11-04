package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"shoplite/internal/models"
	"shoplite/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockOrderService struct{ mock.Mock }

func (m *mockOrderService) Create(input services.CreateOrderInput) (*models.Order, error) {
	args := m.Called(input)
	if v := args.Get(0); v != nil {
		return v.(*models.Order), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *mockOrderService) List() ([]models.Order, error) {
	args := m.Called()
	if v := args.Get(0); v != nil {
		return v.([]models.Order), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *mockOrderService) Get(id uint) (*models.Order, error) {
	args := m.Called(id)
	if v := args.Get(0); v != nil {
		return v.(*models.Order), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestOrderHandler_Create_And_Get(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := new(mockOrderService)
	h := NewOrderHandler(svc)
	r := gin.New()
	r.POST("/orders", h.Create)
	r.GET("/orders", h.List)
	r.GET("/orders/:id", h.GetByID)

	ts := time.Now().UTC().Format(time.RFC3339)
	body := []byte(`{"customer_id":1,"order_date":"` + ts + `","status":"pending","items":[{"product_id":2,"quantity":1,"price":10}]}`)

	svc.On("Create", mock.AnythingOfType("services.CreateOrderInput")).Return(&models.Order{ID: 1, Status: "pending", Items: []models.OrderItem{{ID: 1}}}, nil).Once()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusCreated, w.Code)

	svc.On("List").Return([]models.Order{{ID: 1, Status: "pending"}}, nil).Once()
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest(http.MethodGet, "/orders", nil)
	r.ServeHTTP(w2, req2)
	require.Equal(t, http.StatusOK, w2.Code)

	svc.On("Get", uint(1)).Return(&models.Order{ID: 1}, nil).Once()
	w3 := httptest.NewRecorder()
	req3, _ := http.NewRequest(http.MethodGet, "/orders/1", nil)
	r.ServeHTTP(w3, req3)
	require.Equal(t, http.StatusOK, w3.Code)
}

func TestOrderHandler_Create_Invalid(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := new(mockOrderService)
	h := NewOrderHandler(svc)
	r := gin.New()
	r.POST("/orders", h.Create)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(`{"customer_id":0,"order_date":"not-a-date","items":[]}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusBadRequest, w.Code)
}
