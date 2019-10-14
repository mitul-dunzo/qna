package dtos

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}
