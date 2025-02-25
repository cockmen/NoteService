package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Id       int       `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Created  time.Time `json:"created_at"`
}

type Note struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Created   time.Time `json:"created_at"`
	Updated   time.Time `json:"updated_at"`
	UserEmail string    `json:"user_email"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type QuoteResponse struct {
	Body   string `json:"body"`
	Author string `jsong:"author"`
}
