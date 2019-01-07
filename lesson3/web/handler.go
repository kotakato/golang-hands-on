package web

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/kotakato/golang-hands-on/lesson3/domain"
)

// FilmEchoHandlers は映画関連のHTTPハンドラー。
type FilmEchoHandlers struct {
	repo domain.FilmRepository
}

// SetupFilmEchoHandlers はEchoオブジェクトに映画関連のルーティングを追加する。
func SetupFilmEchoHandlers(e *echo.Echo, repo domain.FilmRepository) {
	h := &FilmEchoHandlers{repo: repo}
	e.GET("/films", h.GetFilms)
	e.GET("/films/:id", h.GetFilm)
}

// GetFilms は映画一覧を取得するハンドラー。
func (h *FilmEchoHandlers) GetFilms(c echo.Context) error {
	films, err := h.repo.GetFilms()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, films)
}

// GetFilm は単一の映画を取得するハンドラー。
func (h *FilmEchoHandlers) GetFilm(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return domain.ErrNotFound
	}
	film, err := h.repo.GetFilm(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, film)
}
