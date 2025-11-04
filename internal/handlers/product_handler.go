package handlers

import (
	"net/http"
	"strconv"

	"shoplite/internal/services"
	"shoplite/internal/utils"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service services.ProductService
}

type CreateProductRequest struct {
	Name  string  `json:"name" binding:"required,min=2,max=160"`
	Price float64 `json:"price" binding:"required,gt=0"`
	Stock int     `json:"stock" binding:"required,gte=0"`
}

func NewProductHandler(s services.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	res, err := h.service.Create(services.CreateProductInput{Name: req.Name, Price: req.Price, Stock: req.Stock})
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.Created(c, "product created", res)
}

func (h *ProductHandler) List(c *gin.Context) {
	res, err := h.service.List()
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.Success(c, "products fetched", res)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
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
	utils.Success(c, "product fetched", res)
}
