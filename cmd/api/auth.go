package main

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Auth struct {
	Issuer        string
	Audience      string
	Secret        string
	TokenExpiry   time.Duration
	RefreshExpire time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}

type jwtUser struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type tokenPairs struct {
	Token        int `json:"token"`
	RefreshToken int `json:"refresh_token"`
}

type Claims struct {
	jwt.RegisteredClaims
}

// func (j *Auth) GenerateToken(user *jwtUser) (tokenPairs, error) {
// 	//Create a token
// 	token :=jwt.New(jwt.SigningMethodES256)  //method can be public key or secret key

// 	//Set Claims
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["name"]= fmt.Sprintf("%s %s", user.FirstName, user.LastName)

// }
