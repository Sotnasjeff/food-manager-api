package db

import (
	"database/sql"

	"github.com/google/uuid"
)

const (
	createMovieQuery   = "INSERT INTO movies (id, name, description) VALUES ($1, $2, $3)"
	findAllMoviesQuery = "SELECT id, name, description FROM movies"
	findMovieById      = "SELECT name, description FROM movies WHERE id = $1"
)

type Movie struct {
	db          *sql.DB
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewMovie(db *sql.DB) *Movie {
	return &Movie{
		db: db,
	}
}

func (m *Movie) CreateMovie(name, description string) (*Movie, error) {
	id := uuid.New().String()
	_, err := m.db.Exec(createMovieQuery, id, name, description)
	if err != nil {
		return nil, err
	}
	return &Movie{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (m *Movie) FindAllMovies() ([]Movie, error) {
	rows, err := m.db.Query(findAllMoviesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	movies := []Movie{}
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		movies = append(movies, Movie{ID: id, Name: name, Description: description})
	}
	return movies, nil
}

func (m *Movie) FindMovieById(id string) (Movie, error) {
	var name, description string
	err := m.db.QueryRow(findMovieById, id).Scan(&name, &description)
	if err != nil {
		return Movie{}, err
	}
	return Movie{ID: id, Name: name, Description: description}, nil
}
