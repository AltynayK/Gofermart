package gofermart

type Orders struct {
	ID         int    `json:"id" db:"id"`
	Number     string `json:"number" db:"number"`
	UserID     int    `json:"user_id" db:"user_id"`
	Status     string `json:"status" db:"status"`
	Accrual    int    `json:"accrual" db:"accrual"`
	UploadedAt string `json:"uploaded_at" db:"uploaded_at"`
	Withdrawn  int    `json:"withdrawn" db:"withdrawn"`
}
type OrdersOut struct {
	Number     string `json:"number" db:"number"`
	Status     string `json:"status" db:"status"`
	Accrual    int    `json:"accrual" db:"accrual"`
	UploadedAt string `json:"uploaded_at" db:"uploaded_at"`
}
type UserBalance struct {
	Current   int `json:"current"`
	Withdrawn int `json:"withdrawn"`
}

type Withdrawals struct {
	Order       string `json:"order"`
	Sum         int    `json:"sum"`
	ProcessedAt string `json:"processed_at"`
}

type OrderBalance struct {
	Order   string `json:"order"`
	Status  string `json:"status"`
	Accrual string `json:"accrual"`
}
