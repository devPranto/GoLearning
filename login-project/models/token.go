package models

import "github.com/dgrijalva/jwt-go/v4"

type JWT struct {
	Header    string `json:"header"`
	Payload   string `json:"payload"`
	Signature string `json:"signature"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
