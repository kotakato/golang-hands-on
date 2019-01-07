package web

import (
	"net/http"

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
}

// GetFilms は映画一覧を取得するハンドラー。
func (h *FilmEchoHandlers) GetFilms(c echo.Context) error {
	films, err := h.repo.GetFilms()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, films)
}
