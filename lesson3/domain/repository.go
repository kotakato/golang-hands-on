package domain

// FilmRepository はFilmのリポジトリ。
type FilmRepository interface {
	GetFilms() ([]*Film, error)
}
