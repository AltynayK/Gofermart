package handler

import (
	"encoding/json"
	"net/http"

	gofermart "github.com/AltynayK/go-musthave-diploma-tpl"
)

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var input gofermart.User
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.services.Authorization.CreateUser(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type loginInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	var input loginInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.services.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
}
