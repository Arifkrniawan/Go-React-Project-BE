package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"movies-be/internal/models"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) AllMovies() ([]*models.Movies, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout) // set timeout kalau user afk 3 detik
	defer cancel()

	sql := `SELECT id, title, release_date, runtime, mpaa_rating, description, coalesce(image, ''), created_at, updated_at
			FROM movies ORDER BY title`

	rows, err := m.DB.QueryContext(ctx, sql)
	if err != nil {
		fmt.Println(rows)
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var movies []*models.Movies

	for rows.Next() {
		var movie models.Movies
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.MPAARATING,
			&movie.Description,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}
	return movies, nil

}

func (m *PostgresDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout) // set timeout kalau user afk 3 detik
	defer cancel()

	sql := `SELECT * FROM users WHERE email = $1`

	var user models.User
	rows := m.DB.QueryRowContext(ctx, sql, email)

	err := rows.Scan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
