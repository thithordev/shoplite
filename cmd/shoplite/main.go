package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"shoplite/config"
	"shoplite/internal/database"
	"shoplite/internal/handlers"
	"shoplite/internal/repositories"
	"shoplite/internal/routes"
	"shoplite/internal/services"
	"shoplite/internal/utils"
)

func main() {
	cfg := config.Load()

	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	v := validator.New()

	// Repositories
	custRepo := repositories.NewCustomerRepository(db)
	prodRepo := repositories.NewProductRepository(db)
	orderRepo := repositories.NewOrderRepository(db)

	// Services
	custService := services.NewCustomerService(custRepo, v)
	prodService := services.NewProductService(prodRepo, v)
	orderService := services.NewOrderService(db, orderRepo, prodRepo, custRepo, v)

	// Handlers
	custHandler := handlers.NewCustomerHandler(custService)
	prodHandler := handlers.NewProductHandler(prodService)
	orderHandler := handlers.NewOrderHandler(orderService)

	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(utils.Logging(), utils.Recovery(), utils.ErrorHandler())

	routes.Register(r, routes.HandlerSet{
		Customers: custHandler,
		Products:  prodHandler,
		Orders:    orderHandler,
	})

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("ShopLite API running on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
