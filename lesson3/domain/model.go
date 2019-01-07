package domain

// Film は映画を表す。単純にするため一部のカラムのみ。
type Film struct {
	FilmID      int    `json:"filmId"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	ReleaseYear int    `json:"releaseYear" validate:"required"`
	LanguageID  int    `json:"languageId" validate:"required"`
}
