package model

import "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	ID             int64
	Username       string
	ProfilePicture string
	jwt.StandardClaims
}
