package handler

import (
	"io/ioutil"
	"net/http"

	gofermart "github.com/AltynayK/go-musthave-diploma-tpl"
	"github.com/gin-gonic/gin"
)

func (h *Handler) loadingOrders(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}
	body := c.Request.Body
	input, _ := ioutil.ReadAll(body)
	id, err := h.services.Order.Create(userID, string(input))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type getAllOrdersResponse struct {
	Data []gofermart.OrdersOut `json:"data"`
}

func (h *Handler) receivingOrders(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}
	orders, err := h.services.Order.GetAll(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllOrdersResponse{
		Data: orders,
	})
}
func (h *Handler) receivingBalance(c *gin.Context) {

}
func (h *Handler) withdrawBalance(c *gin.Context) {

}
func (h *Handler) withdrawBalanceHistory(c *gin.Context) {

}
