package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Auth struct {
	Issuer        string
	Audience      string
	Secret        string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}

type jwtUser struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type TokenPairs struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	jwt.RegisteredClaims
}

func (j *Auth) GenerateToken(user *jwtUser) (TokenPairs, error) {
	//Create a token
	token := jwt.New(jwt.SigningMethodHS256) //method can be public key or secret key

	//Set Claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	claims["sub"] = fmt.Sprint(user.Id)
	claims["aud"] = j.Audience
	claims["iss"] = j.Issuer
	claims["iat"] = time.Now().UTC().Unix()
	claims["typ"] = "JWT"

	//Set Expiry for JWT
	claims["exp"] = time.Now().Add(j.TokenExpiry).Unix()

	//Create a signed token
	signedAccessToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	//Create a refresh token and set claims
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["sub"] = fmt.Sprint(user.Id)
	refreshTokenClaims["iat"] = time.Now().UTC().Unix()

	//Set token refresh expiry
	refreshTokenClaims["exp"] = time.Now().Add(j.RefreshExpiry).Unix()

	//Create signed refersh token
	signedRefreshToken, err := refreshToken.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	//Create TokenPairs and populate with sign tokens
	var tokenPairs = TokenPairs{
		Token:        signedAccessToken,
		RefreshToken: signedRefreshToken,
	}

	//return token
	return tokenPairs, nil
}

func (j *Auth) GetRefreshCookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Name: j.CookieName,
		Path: j.CookiePath,
		Value: refreshToken,
		Domain: j.CookieDomain,
		Expires: time.Now().Add(j.RefreshExpiry),
		MaxAge: int(j.RefreshExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Secure: true,
	}
}

func (j *Auth) GetExpiredRefreshCookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Name: j.CookieName,
		Path: j.CookiePath,
		Value: "",
		Domain: j.CookieDomain,
		Expires: time.Unix(0,0),
		MaxAge: -1,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Secure: true,
	}
}