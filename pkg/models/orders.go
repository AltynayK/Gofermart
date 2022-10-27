package models

type Orders struct {
	ID         int     `json:"id" db:"id"`
	Number     string  `json:"number" db:"number"`
	UserID     int     `json:"user_id" db:"user_id"`
	Status     string  `json:"status" db:"status"`
	Accrual    float32 `json:"accrual" db:"accrual"`
	UploadedAt string  `json:"uploaded_at" db:"uploaded_at"`
	Withdrawn  float32 `json:"withdrawn" db:"withdrawn"`
}
type OrdersOut struct {
	Number     string  `json:"number" db:"number"`
	Status     string  `json:"status" db:"status"`
	Accrual    float32 `json:"accrual" db:"accrual"`
	UploadedAt string  `json:"uploaded_at" db:"uploaded_at"`
}
type UserBalance struct {
	Current   float32 `json:"current"`
	Withdrawn float32 `json:"withdrawn"`
}

type Withdrawals struct {
	Order       string  `json:"order" db:"number"`
	Sum         float32 `json:"sum" db:"withdrawn"`
	ProcessedAt string  `json:"processed_at" db:"processed_at"`
}
type OrderBalance struct {
	Order   string  `json:"order" `
	Status  string  `json:"status"`
	Accrual float32 `json:"accrual"`
}
