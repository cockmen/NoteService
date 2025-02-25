package users

import "time"

type User struct {
	Id       int       `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Created  time.Time `json:"created_at"`
}
