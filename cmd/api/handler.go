package main

import (
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
