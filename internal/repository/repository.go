package repository

import (
	"database/sql"
	"movies-be/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movies, error)
}
