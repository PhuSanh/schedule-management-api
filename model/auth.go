package model

import "github.com/dgrijalva/jwt-go"

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthClaims struct {
	Id 			int64 	`json:"id"`
	Username 	string 	`json:"username"`
	Email 		string 	`json:"email"`
	jwt.StandardClaims
}

type AuthResponse struct {
	Token 		string 	`json:"token"`
	Username 	string 	`json:"username"`
}