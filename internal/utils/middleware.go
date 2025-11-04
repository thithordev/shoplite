package utils

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logging() gin.HandlerFunc {
	return gin.Logger()
}

func Recovery() gin.HandlerFunc {
	return gin.Recovery()
}

// ErrorHandler ensures errors are returned in a standard JSON format if any unhandled error bubbles up.
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			log.Println("request error:", c.Errors.String())
			Error(c, http.StatusInternalServerError, "internal server error", nil)
		}
	}
}
