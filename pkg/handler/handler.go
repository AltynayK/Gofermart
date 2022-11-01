package handler

import (
	"github.com/AltynayK/go-musthave-diploma-tpl/configs"
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/repository"
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const chanVal = 5

type Handler struct {
	config          *configs.Config
	services        *service.Service
	db              *sqlx.DB
	repos           *repository.Repository
	queueForAccrual chan string
}

func NewHandler() *Handler {
	config := configs.NewConfig()
	db := repository.NewPostgresDB(config)
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	return &Handler{
		config:          config,
		db:              db,
		repos:           repos,
		services:        services,
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
