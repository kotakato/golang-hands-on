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
	e.POST("/films", h.CreateFilm)
	e.DELETE("/films/:id", h.DeleteFilm)
	e.PUT("/films/:id", h.UpdateFilm)
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

// CreateFilm は映画を新規作成するハンドラー。
func (h *FilmEchoHandlers) CreateFilm(c echo.Context) error {
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return echo.NewHTTPError(http.StatusUnsupportedMediaType, "Content-Type must be application/json")
	}
	var film domain.Film
	if err := c.Bind(&film); err != nil {
		return err
	}
	if err := c.Validate(film); err != nil {
		return err
	}
	f, err := h.repo.InsertFilm(&film)
	if err != nil {
		return err
	}
	c.Response().Header().Set("Location", c.Echo().URI(h.GetFilm, f.FilmID))
	return c.JSON(http.StatusCreated, f)
}

// DeleteFilm は映画を削除するハンドラー。
func (h *FilmEchoHandlers) DeleteFilm(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return domain.ErrNotFound
	}
	err = h.repo.DeleteFilm(id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *FilmEchoHandlers) UpdateFilm(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return domain.ErrNotFound
	}
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return echo.NewHTTPError(http.StatusUnsupportedMediaType, "Content-Type must be application/json")
	}

	var film domain.Film
	if err := c.Bind(&film); err != nil {
		return err
	}
	if err := c.Validate(film); err != nil {
		return err
	}
	f, err := h.repo.UpdateFilm(id, &film)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, f)
}
