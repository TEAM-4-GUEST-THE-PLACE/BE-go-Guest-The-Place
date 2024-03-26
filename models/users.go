package models

import (
	"time"
)

type Users struct {
    ID              uint     `json:"id" gorm:"primaryKey"`
    Diamonds_totals int      `json:"diamonds_totals"`
    Fullname        string   `json:"fullname"`
    Username        string   `json:"username"`
    Email           string   `json:"email"`
    Avatars         uint  	 
    Created_at      time.Time
    Updated_at      time.Time
}


type UpdateUserRequest struct {
	Fullname 		string `json:"fullname" form:"fullname"`
	Username 		string `json:"username" form:"username"`
	Email  			string `json:"email" form:"email"`
}


type CreateUserRequest struct {
	ID              uint   `json:"id" gorm:"primaryKey" form:"id" `
	Diamonds_totals int    `json:"diamonds_totals" form:"diamonds_totals" `
	Fullname        string `json:"fullname" form:"fullname" validate:"required"`
	Username        string `json:"username" form:"username" validate:"required"`
	Email           string `json:"email" form:"email" validate:"required"`
	Avatars         uint

	Created_at time.Time
	Updated_at time.Time

}