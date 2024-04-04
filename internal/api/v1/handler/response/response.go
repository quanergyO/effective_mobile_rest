package response

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

func NewError(c *gin.Context, statusCode int, message string) {
	slog.Error("Error message:", message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

type errorResponse struct {
	Message string `json:"message"`
}
