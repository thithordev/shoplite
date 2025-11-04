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

type mockProductService struct{ mock.Mock }

func (m *mockProductService) Create(input services.CreateProductInput) (*models.Product, error) {
	args := m.Called(input)
	if v := args.Get(0); v != nil {
		return v.(*models.Product), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *mockProductService) List() ([]models.Product, error) {
	args := m.Called()
	if v := args.Get(0); v != nil {
		return v.([]models.Product), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *mockProductService) Get(id uint) (*models.Product, error) {
	args := m.Called(id)
	if v := args.Get(0); v != nil {
		return v.(*models.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestProductHandler_Create_And_List(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := new(mockProductService)
	h := NewProductHandler(svc)
	r := gin.New()
	r.POST("/products", h.Create)
	r.GET("/products", h.List)

	svc.On("Create", services.CreateProductInput{Name: "Widget", Price: 2.5, Stock: 3}).Return(&models.Product{ID: 1, Name: "Widget", Price: 2.5, Stock: 3}, nil).Once()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(`{"name":"Widget","price":2.5,"stock":3}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusCreated, w.Code)

	svc.On("List").Return([]models.Product{{ID: 1, Name: "Widget", Price: 2.5, Stock: 3}}, nil).Once()
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest(http.MethodGet, "/products", nil)
	r.ServeHTTP(w2, req2)
	require.Equal(t, http.StatusOK, w2.Code)
}

func TestProductHandler_Create_Invalid(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := new(mockProductService)
	h := NewProductHandler(svc)
	r := gin.New()
	r.POST("/products", h.Create)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(`{"name":"W","price":-1,"stock":-1}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusBadRequest, w.Code)
}
