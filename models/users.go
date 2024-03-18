package models

import (
	"time"
)

type Users struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Question_id    int    `json:"question_id"`
	Avatar_id      int    `json:"avatar_id"`
	Diamond_totals int    `json:"diamond_totals"`
	Fullname       string `json:"fullname"`
	Username       string `json:"username"`
	Email          string `json:"email"`

	Created_at time.Time
	Updated_at time.Time
}

// type UpdateUser struct {
// 	ID             uint   `json:"id" gorm:"primaryKey"`
// 	Question_id    int    `json:"question_id"`
// 	Avatar_id      int    `json:"avatar_id"`
// 	Diamond_totals int    `json:"diamond_totals"`

// 	Fullname string `json:"fullname" form:"fullname"`
// 	Username string `json:"username" form:"username"`
// 	Email    string `json:"email" form:"email"`

// 	Created_at time.Time
// 	Updated_at time.Time
// }
