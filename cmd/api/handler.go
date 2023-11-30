package main

import (
	"log"
	"net/http"
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

func (app *application) authenticate(w http.ResponseWriter, r *http.Request){
	//Read json payload
	var requestPayload struct {
		Email string `json:"email"`
		Password string `json:"password`
	}

	err := app.readJSON(w, r, requestPayload)
	if err != nil {
		app.errJSON(w, err)
	}
	//Validate user against database
	user, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		app.errJSON(w, err, http.StatusBadRequest)
		return
	}

	//Check password
	valid, err := user.PasswordMatch(user.Password)
	if err != nil || !valid {
		app.errJSON(w, err, http.StatusBadRequest)
		return
	}

	
	//Create JWT user
	userr := jwtUser{
		Id: user.ID,
		FirstName: user.FirstName,
		LastName: user.LastName,
	}

	token, err := app.auth.GenerateToken(&userr)
	if err != nil {
		app.errJSON(w, err)
		return
	}

	refreshedCookie := app.auth.GetRefreshCookie(token.RefreshToken)
	http.SetCookie(w, refreshedCookie)

	app.writeJson(w, http.StatusAccepted, token)
}