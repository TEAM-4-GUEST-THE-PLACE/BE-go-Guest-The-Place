package models

type LoginRequest struct {
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
}

type LoginResponse struct {
	Fullname       string `json:"fullname"`
	Email          string `json:"email"`
	Diamond_totals int    `json:"diamond_totals" form:"diamond_totals" `
	Avatar_id      int    `json:"avatar_id" form:"avatar_id" `
}
