package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func main() {
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)

}
