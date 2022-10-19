package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	fmt.Print(message)
	c.AbortWithStatusJSON(statusCode, error{message})
}
