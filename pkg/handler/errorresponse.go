package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	fmt.Print(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
