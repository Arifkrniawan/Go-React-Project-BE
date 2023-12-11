package dbrepo

import (
	"context"
	"database/sql"
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

	sql := `select id, email, first_name, last_name, password,
	created_at, updated_at from users where email = $1`

	var user models.User
	rows := m.DB.QueryRowContext(ctx, sql, email)

	err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *PostgresDBRepo) GetUserById(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout) // set timeout kalau user afk 3 detik
	defer cancel()

	sql := `select id, email, first_name, last_name, password,
	created_at, updated_at from users where id = $1`

	var user models.User
	rows := m.DB.QueryRowContext(ctx, sql, id)

	err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}