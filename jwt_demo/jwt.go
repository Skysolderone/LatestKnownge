package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secrecet = []byte("testjwt")

type myclaims struct {
	Foo string
	jwt.RegisteredClaims
}

func main() {
	claims := myclaims{"testjwt", jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		Issuer:    "test",
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secrecet)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ss)
	parsetoken, err := jwt.ParseWithClaims(ss, &myclaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("testjwt"), nil
	})
	if claims, ok := parsetoken.Claims.(*myclaims); ok {
		fmt.Println(claims.Foo)
	} else {
		fmt.Println(err)
	}
	switch {
	case parsetoken.Valid:
		fmt.Println("You look nice today")
	case errors.Is(err, jwt.ErrTokenMalformed):
		fmt.Println("That's not even a token")
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		// Invalid signature
		fmt.Println("Invalid signature")
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		// Token is either expired or not active yet
		fmt.Println("Timing is everything")
	default:
		fmt.Println("Couldn't handle this token:", err)
	}
}
