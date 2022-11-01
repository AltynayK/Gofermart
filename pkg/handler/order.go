package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/models"
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/service"
	"github.com/gin-gonic/gin"
)

func (h *Handler) loadingOrders(c *gin.Context) {
	c.Set("content-type", "plain/text")
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
	if !service.ValidByLuhn(num) {
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
	h.WriteOrderToChan(string(input))
	c.AbortWithStatus(http.StatusAccepted)
}
func (h *Handler) WriteOrderToChan(processingOrder string) {
	var s *Server
	s.queueForAccrual <- processingOrder

}

func (h *Handler) GetOrderAccrual() {
	var orderNumber string
	var s *Server
	for i := range s.queueForAccrual {
		orderNumber = i
		var datas models.OrderBalance
		resp, err := http.Get(NewServer().config.AccrualSystemAddress + "/api/orders/" + orderNumber)
		if err != nil {
			fmt.Print(err)
		}
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Print(err)
		}
		resp.Body.Close()
		err = json.Unmarshal(responseBody, &datas)
		if err != nil {
			fmt.Print(err)
		}
		_, err = h.services.Order.PostBalance(datas)
		if err != nil {
			fmt.Print(err)
			return
		}

		userID, err := h.services.Order.GetOrderUserID(orderNumber)
		//Взаимодействие с системой расчёта начислений баллов лояльности
		if err != nil {
			fmt.Print(err)
			return
		}
		current, err := h.services.Order.GetUserCurrent(userID)
		if err != nil {
			return
		}
		newcurrent := current + datas.Accrual

		_, err = h.services.Order.UpdateUserBalance(userID, newcurrent)
		if err != nil {
			fmt.Print(err)
			return
		}

	}

}

func (h *Handler) receivingOrders(c *gin.Context) {
	c.Set("Content-type", "application/json")
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
	c.JSON(http.StatusOK, orders)
}

func (h *Handler) receivingBalance(c *gin.Context) {
	c.Set("content-type", "application/json")
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	current, err := h.services.Order.GetUserCurrent(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	withdrawn, err := h.services.Order.GetUserWithdrawn(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.UserBalance{
		Current:   current,
		Withdrawn: withdrawn,
	})
}

func (h *Handler) withdrawBalance(c *gin.Context) {
	c.Set("content-type", "application/json")

	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var input models.Withdrawals
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//проверка на корректность ввода с помощью алгоритма Луна
	num, _ := strconv.Atoi(string(input.Order))
	if !service.ValidByLuhn(num) {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}
	current, err := h.services.Order.GetUserCurrent(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	//код 402, на счету недостаточно средств
	if current < float32(input.Sum) {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	//проверка номера заказа на существование
	order, err := h.services.Order.GetOrder(num)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if order == nil {
		err = h.services.Order.PostNewWithdrawBalance(input, userID)
		if err != nil {
			newErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
			return
		}
	} else {
		_, err = h.services.Order.PostWithdrawBalance(input)
		if err != nil {
			newErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
			return
		}
	}
	newcurrent := current - float32(input.Sum)
	h.services.Order.UpdateUserBalance(userID, newcurrent)
	c.AbortWithStatus(http.StatusOK)
}

func (h *Handler) withdrawBalanceHistory(c *gin.Context) {
	c.Set("content-type", "application/json")

	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	withdrawals, err := h.services.Order.GetAllWithdrawals(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if withdrawals == nil {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.JSON(http.StatusOK, withdrawals)
}
