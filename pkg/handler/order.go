package handler

import (
	"io/ioutil"
	"net/http"
	"strconv"

	gofermart "github.com/AltynayK/go-musthave-diploma-tpl"
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/service"
	"github.com/gin-gonic/gin"
)

func (h *Handler) loadingOrders(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	body := c.Request.Body
	input, _ := ioutil.ReadAll(body)
	//проверка на корректность ввода с помощью алгоритма Луна
	num, _ := strconv.Atoi(string(input))
	if !service.Valid(num) {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	//создание нового заказа
	err = h.services.Order.Create(userID, string(input))
	//проверка код ответа 200, номер заказа уже был загружен этим пользователем
	if err != nil && err.Error() == "pq: duplicate key value violates unique constraint \"orders_number_key\"" {
		c.AbortWithStatus(http.StatusOK)
		return
	}
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.AbortWithStatus(http.StatusAccepted)

}

type getAllOrdersResponse struct {
	Data []gofermart.OrdersOut `json:"data"`
}

func (h *Handler) receivingOrders(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	orders, err := h.services.Order.GetAll(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if orders == nil {
		c.AbortWithStatus(http.StatusNoContent)
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
