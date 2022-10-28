package handler

import (
	"github.com/AltynayK/go-musthave-diploma-tpl/configs"
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/repository"
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	config          *configs.Config
	services        *service.Service
	db              *sqlx.DB
	repos           *repository.Repository
	queueForAccrual chan string
}

const chanVal = 5

func NewHandler() *Handler {
	return &Handler{
		config:          configs.NewConfig(),
		db:              repository.NewPostgresDB(configs.NewConfig()),
		repos:           repository.NewRepository(repository.NewPostgresDB(configs.NewConfig())),
		services:        service.NewService(repository.NewRepository(repository.NewPostgresDB(configs.NewConfig()))),
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
