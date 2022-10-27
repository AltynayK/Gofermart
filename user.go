package gofermart

type User struct {
	ID       int     `json:"-" db:"id"`
	Login    string  `json:"login" binding:"required"`
	Password string  `json:"password" binding:"required"`
	Current  float32 `json:"-"`
}
