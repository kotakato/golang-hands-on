package domain

// FilmRepository はFilmのリポジトリ。
type FilmRepository interface {
	GetFilms() ([]*Film, error)
	GetFilm(id int) (*Film, error)
	InsertFilm(film *Film) (*Film, error)
	DeleteFilm(id int) error
	UpdateFilm(id int, film *Film) (*Film, error)
}
