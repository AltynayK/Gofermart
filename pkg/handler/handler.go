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
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	auth := router.Group("/api/user")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.login)

	}
	api := router.Group("/api/user", h.userIdentity)
	{
		api.POST("/orders", h.loadingOrders)
		api.GET("/orders", h.receivingOrders)
		api.GET("/balance", h.receivingBalance)
		api.POST("/balance/withdraw", h.withdrawBalance)
		api.GET("/withdrawals", h.withdrawBalanceHistory)

	}

	return router
}
