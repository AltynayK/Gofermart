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
	//проверка код ответа 400, неверный формат запроса
	input, err := ioutil.ReadAll(body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	//проверка на корректность ввода с помощью алгоритма Луна
	num, _ := strconv.Atoi(string(input))
	if !service.Valid(num) {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}
	//проверка код ответа 200, номер заказа уже был загружен этим пользователем
	orders, err := h.services.Order.GetOrderByUserAndNumber(userID, num)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if orders != nil {
		c.AbortWithStatus(http.StatusOK)
		return
	}
	//проверка код ответа 409, номер заказа уже был загружен другим пользователем
	order, err := h.services.Order.GetOrder(num)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if order != nil {
		c.AbortWithStatus(http.StatusConflict)
		return
	}
	//создание нового заказа
	err = h.services.Order.Create(userID, string(input))
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
