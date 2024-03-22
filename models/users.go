package models

import (
	"time"
)

type Users struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Avatar_id      int    `json:"avatar_id"`
	Diamond_totals int    `json:"diamond_totals"`
	Fullname       string `json:"fullname"`
	Username       string `json:"username"`
	Email          string `json:"email"`

	Created_at time.Time
	Updated_at time.Time
}


type UpdateUserRequest struct {
	Fullname 		string `json:"fullname" form:"fullname"`
	Username 		string `json:"username" form:"username"`
	Email  			string `json:"email" form:"email"`
}


type CreateUserRequest struct {
	ID             uint   `json:"id" gorm:"primaryKey" form:"id" `
	Avatar_id      int    `json:"avatar_id" form:"avatar_id" `
	Diamond_totals int    `json:"diamond_totals" form:"diamond_totals" `
	Fullname       string `json:"fullname" form:"fullname" validate:"required"`
	Username       string `json:"username" form:"username" validate:"required"`
	Email          string `json:"email" form:"email" validate:"required"`

	Created_at time.Time
	Updated_at time.Time

}