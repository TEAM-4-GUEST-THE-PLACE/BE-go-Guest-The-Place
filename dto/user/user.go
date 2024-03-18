package usersdto


import (
	"time"
)

type UpdateUserRequest struct {
	Fullname 		string `json:"fullname" form:"fullname"`
	Username 		string `json:"username" form:"username"`
	Email  			string `json:"email" form:"email"`
	
	Created_at 		time.Time
	Updated_at 		time.Time

}
