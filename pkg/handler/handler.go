package handler

import (
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api/user")
	{
		api.POST("/register", h.register)
		api.POST("/login", h.login)
		api.POST("/orders", h.loadingOrders)
		api.GET("/orders", h.receivingOrders)
		api.GET("/balance", h.receivingBalance)
		api.POST("/balance/withdraw", h.withdrawBalance)
		api.GET("/balance/withdrawals", h.withdrawBalanceHistory)

	}
	return router
}
