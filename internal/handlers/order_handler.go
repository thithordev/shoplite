package handlers

import (
	"net/http"
	"strconv"
	"time"

	"shoplite/internal/services"
	"shoplite/internal/utils"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service services.OrderService
}

type OrderItemRequest struct {
	ProductID uint    `json:"product_id" binding:"required,gt=0"`
	Quantity  int     `json:"quantity" binding:"required,gt=0"`
	Price     float64 `json:"price" binding:"required,gt=0"`
}

type CreateOrderRequest struct {
	CustomerID uint               `json:"customer_id" binding:"required,gt=0"`
	OrderDate  time.Time          `json:"order_date" binding:"required"`
	Status     string             `json:"status" binding:"omitempty,oneof=pending paid shipped cancelled"`
	Items      []OrderItemRequest `json:"items" binding:"required,min=1,dive"`
}

func NewOrderHandler(s services.OrderService) *OrderHandler {
	return &OrderHandler{service: s}
}

func (h *OrderHandler) Create(c *gin.Context) {
	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	items := make([]services.OrderItemInput, 0, len(req.Items))
	for _, it := range req.Items {
		items = append(items, services.OrderItemInput{ProductID: it.ProductID, Quantity: it.Quantity, Price: it.Price})
	}
	res, err := h.service.Create(services.CreateOrderInput{
		CustomerID: req.CustomerID,
		OrderDate:  req.OrderDate,
		Status:     req.Status,
		Items:      items,
	})
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.Created(c, "order created", res)
}

func (h *OrderHandler) List(c *gin.Context) {
	res, err := h.service.List()
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.Success(c, "orders fetched", res)
}

func (h *OrderHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid id", nil)
		return
	}
	res, err := h.service.Get(uint(id))
	if err != nil {
		utils.Error(c, http.StatusNotFound, err.Error(), nil)
		return
	}
	utils.Success(c, "order fetched", res)
}
