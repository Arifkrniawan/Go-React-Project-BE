package main

import (
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var g = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "Hai",
		Message: "Ok",
		Version: "1.0.0",
	}

	_ = app.writeJson(w, http.StatusOK, g)
}

func (app *application) Movies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()
	if err != nil {
		app.errJSON(w, err)
		return
	}

	_ = app.writeJson(w, http.StatusOK, movies)
}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	//Read json payload
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	log.Print(requestPayload.Email, requestPayload.Password)

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errJSON(w, err, http.StatusBadRequest)
	}

	//Validate user against database
	user, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		app.errJSON(w, err, http.StatusBadRequest)
		return
	}
	log.Print(user)
	//Check password
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errJSON(w, err, http.StatusBadRequest)
		log.Print(requestPayload.Password)
		return
	}

	//Create JWT user
	userr := jwtUser{
		Id:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	token, err := app.auth.GenerateTokenPair(&userr)
	if err != nil {
		app.errJSON(w, err)
		return
	}

	refreshedCookie := app.auth.GetRefreshCookie(token.RefreshToken)
	http.SetCookie(w, refreshedCookie)

	app.writeJson(w, http.StatusAccepted, token)
}

// func (app *application) refreshToken(w http.ResponseWriter, r *http.Request) {
// 	for _, cookie := range r.Cookies() {
// 		if cookie.Name == app.auth.CookieName {
// 			claims := &Claims{}
// 			refreshToken := cookie.Value

// 			_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error){
// 				return []byte(app.JWTSecret), nil
// 			})

// 			if err != nil {
// 				app.errJSON(w, err, http.StatusUnauthorized)
// 				return
// 			}

// 			userID, err := strconv.Atoi(claims.Subject)
// 			if err != nil {
// 				app.errJSON(w, err, http.StatusUnauthorized)
// 				return
// 			}

// 			user, err := app.DB.GetUserByEmail()
// 		}
// 	}
// }