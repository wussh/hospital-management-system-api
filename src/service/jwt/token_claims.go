package jwt

import "github.com/golang-jwt/jwt"

type Claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
