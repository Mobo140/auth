package model

import jwt "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}
