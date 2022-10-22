package handler

import (
	"net/http"

	gofermart "github.com/AltynayK/go-musthave-diploma-tpl"
	"github.com/gin-gonic/gin"
)

func (h *Handler) register(c *gin.Context) {
	c.Set("content-type", "application/json")
	var input gofermart.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusConflict, err.Error())
		return
	}
	token, err := h.services.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusConflict, err.Error())
		return
	}
	c.Header("Authorization", token)
	c.JSON(http.StatusOK, input)
}

type loginInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) login(c *gin.Context) {
	c.Set("content-type", "application/json")
	var input loginInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.services.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusConflict, err.Error())
		return
	}
	c.Header("Authorization", token)
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
