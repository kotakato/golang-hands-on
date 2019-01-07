package web

import (
	"net/http"

	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/kotakato/golang-hands-on/lesson3/domain"
)

// EchoCustomHTTPErrorHandler はエラーをレスポンスに変換する。
func EchoCustomHTTPErrorHandler(err error, c echo.Context) {
	err = errorResponse(err, c)
	c.Echo().DefaultHTTPErrorHandler(err, c)
}

func errorResponse(err error, c echo.Context) error {
	switch err.(type) {
	case *echo.HTTPError:
		return err
	case validator.ValidationErrors:
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
	}

	switch err {
	case domain.ErrNotFound:
		return echo.NewHTTPError(http.StatusNotFound, err.Error()).SetInternal(err)
	case domain.ErrConflict:
		return echo.NewHTTPError(http.StatusConflict, err.Error()).SetInternal(err)
	}
	c.Logger().Error(err)
	return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
}
