package routes

import (
	"shoplite/internal/handlers"

	"github.com/gin-gonic/gin"
)

type HandlerSet struct {
	Customers *handlers.CustomerHandler
	Products  *handlers.ProductHandler
	Orders    *handlers.OrderHandler
}

func Register(r *gin.Engine, h HandlerSet) {
	r.POST("/customers", h.Customers.Create)
	r.GET("/customers", h.Customers.List)
	r.GET("/customers/:id", h.Customers.GetByID)

	r.POST("/products", h.Products.Create)
	r.GET("/products", h.Products.List)
	r.GET("/products/:id", h.Products.GetByID)

	r.POST("/orders", h.Orders.Create)
	r.GET("/orders", h.Orders.List)
	r.GET("/orders/:id", h.Orders.GetByID)
}
