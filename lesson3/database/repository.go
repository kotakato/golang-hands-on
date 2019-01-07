package database

import (
	"database/sql"

	"github.com/kotakato/golang-hands-on/lesson3/domain"
)

// FilmSQLRepository はFilmRepositoryのSQLによる実装。
type FilmSQLRepository struct {
	db *sql.DB
}

// NewFilmSQLRepository はFilmSQLRepositoryを作成する。
func NewFilmSQLRepository(db *sql.DB) domain.FilmRepository {
	return &FilmSQLRepository{db: db}
}

// GetFilms はすべてのFilmを取得する。
func (r *FilmSQLRepository) GetFilms() ([]*domain.Film, error) {
	rows, err := r.db.Query(`
		SELECT film_id, title, description, release_year, language_id
		FROM film`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	films, err := scanRows(rows)
	if err != nil {
		return nil, err
	}
	return films, nil
}

func scanRows(rows *sql.Rows) ([]*domain.Film, error) {
	films := make([]*domain.Film, 0)
	for rows.Next() {
		var film domain.Film
		err := rows.Scan(
			&film.FilmID,
			&film.Title,
			&film.Description,
			&film.ReleaseYear,
			&film.LanguageID,
		)
		if err != nil {
			return nil, err
		}
		films = append(films, &film)
	}
	return films, nil
}
