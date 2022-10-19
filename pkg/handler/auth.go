package handler

import (
	"net/http"

	gofermart "github.com/AltynayK/go-musthave-diploma-tpl"
	"github.com/gin-gonic/gin"
)

func (h *Handler) register(c *gin.Context) {
	var input gofermart.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) login(c *gin.Context) {

}
