package models

import (
	"time"
)

type User struct {
	ID            	 uint   `json:"id" gorm:"primaryKey"`
	Question_id  	 int	 `json:"question_id"`
	Avatar_id      	 int	 `json:"avatar_id"`
	Diamond_totals	 int	 `json:"diamond_totals"`
	Fullname      	 string `json:"fullname"`
	Username      	 string `json:"username"`
	Email         	 string `json:"email"`

	Created_at		time.Time
	Updated_at		time.Time
}
