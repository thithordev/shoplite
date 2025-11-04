package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Response{
		Status:  "error",
		Message: message,
		Data:    data,
	})
}
