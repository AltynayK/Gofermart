package handler

import (
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/user/register", h.register).Methods("POST")
	router.HandleFunc("/api/user/login", h.login).Methods("POST")
	router.HandleFunc("/api/user/orders", h.loadingOrders).Methods("POST")
	router.HandleFunc("/api/user/orders", h.receivingOrders).Methods("GET")
	router.HandleFunc("/api/user/balance", h.receivingBalance).Methods("GET")
	router.HandleFunc("/api/user/balance/withdraw", h.withdrawBalance).Methods("POST")
	router.HandleFunc("/api/user/balance/withdrawals", h.withdrawBalanceHistory).Methods("GET")

	return router
}
