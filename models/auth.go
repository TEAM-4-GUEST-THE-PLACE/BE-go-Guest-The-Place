package models

type LoginRequest struct {
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
}

type LoginResponse struct {
	Fullname       	string `json:"fullname"`
	Email          	string `json:"email"`
	Diamonds_totals int    `json:"diamonds_totals" form:"diamonds_totals" `
}
