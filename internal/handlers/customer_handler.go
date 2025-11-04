package handlers

import (
	"net/http"
	"strconv"

	"shoplite/internal/services"
	"shoplite/internal/utils"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	service services.CustomerService
}

type CreateCustomerRequest struct {
	Name  string `json:"name" binding:"required,min=2,max=120"`
	Email string `json:"email" binding:"required,email"`
}

func NewCustomerHandler(s services.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: s}
}

func (h *CustomerHandler) Create(c *gin.Context) {
	var req CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	res, err := h.service.Create(services.CreateCustomerInput{Name: req.Name, Email: req.Email})
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.Created(c, "customer created", res)
}

func (h *CustomerHandler) List(c *gin.Context) {
	res, err := h.service.List()
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.Success(c, "customers fetched", res)
}

func (h *CustomerHandler) GetByID(c *gin.Context) {
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
	utils.Success(c, "customer fetched", res)
}
