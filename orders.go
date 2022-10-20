package gofermart

type Orders struct {
	ID         int    `json:"id"`
	Number     string `json:"number"`
	UserID     int    `json:"user_id"`
	Status     string `json:"status"`
	Accrual    int    `json:"accrual"`
	UploadedAt string `json:"uploaded_at"`
	Withdrawn  int    `json:"withdrawn"`
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
