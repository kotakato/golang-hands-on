package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"

	"github.com/kotakato/golang-hands-on/lesson3/database"
	"github.com/kotakato/golang-hands-on/lesson3/web"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Validator = web.NewEchoCustomValidator()
	e.HTTPErrorHandler = web.EchoCustomHTTPErrorHandler
	e.Use(middleware.Logger())

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:password@localhost/postgres?sslmode=disable"
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		e.Logger.Fatal(err)
	}

	repo := database.NewFilmSQLRepository(db)
	web.SetupFilmEchoHandlers(e, repo)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
