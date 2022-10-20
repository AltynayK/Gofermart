package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) loadingOrders(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	body := c.Request.Body
	input, _ := ioutil.ReadAll(body)
	id, err := h.services.Order.Create(userId, string(input))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}
func (h *Handler) receivingOrders(c *gin.Context) {

}
func (h *Handler) receivingBalance(c *gin.Context) {

}
func (h *Handler) withdrawBalance(c *gin.Context) {

}
func (h *Handler) withdrawBalanceHistory(c *gin.Context) {

}
