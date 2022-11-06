package handler

import (
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/service"
	"github.com/gin-gonic/gin"
)

const chanVal = 5

type Handler struct {
	//services *service.Service
	auth            *service.AuthService
	order           *service.OrderService
	queueForAccrual chan string
}

func NewHandler() *Handler {

	//services := service.NewService(NewServer().repos)
	auth := service.NewAuthService(NewServer().repos)
	order := service.NewOrderService(NewServer().repos)
	return &Handler{
		//services: services,
		auth:            auth,
		order:           order,
		queueForAccrual: make(chan string, chanVal),
	}
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
	go h.GetOrderAccrual()
	return router
}
