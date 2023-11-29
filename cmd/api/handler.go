package main

import (
	"encoding/json"
	"fmt"
	"movies-be/internal/models"
	"net/http"
	"time"
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

	out, err := json.Marshal(g)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) Movies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movies

	tp, _ := time.Parse("01-12-2012", "04-03-2001")

	Spiderman1 := models.Movies{
		ID:          1,
		Title:       "Spiderman 1",
		ReleaseDate: tp,
		Runtime:     120,
		MPAARATING:  "R",
		Description: "Good Movies!",
	}

	movies = append(movies, Spiderman1)

	tp, _ = time.Parse("01-12-2012", "25-02-2003")

	Spiderman2 := models.Movies{
		ID:          2,
		Title:       "Spiderman 2",
		ReleaseDate: tp,
		Runtime:     120,
		MPAARATING:  "R",
		Description: "Another Good Movies!",
	}

	movies = append(movies, Spiderman2)

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
