package handler

import (
	"net/http"
)

func (h *Handler) loadingOrders(w http.ResponseWriter, r *http.Request) {
	// userID, err := getUserID(r)
	// if err != nil {
	// 	return
	// }
	// body := r.Request.Body
	// input, _ := ioutil.ReadAll(body)
	// h.services.Order.Create(userID, string(input))
	// // if err != nil {
	// // 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// // 	return
	// // }
	// w.Writer.Status(http.StatusOK)

}
func (h *Handler) receivingOrders(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) receivingBalance(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) withdrawBalance(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) withdrawBalanceHistory(w http.ResponseWriter, r *http.Request) {

}
