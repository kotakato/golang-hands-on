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
	// Echoの設定
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Validator = web.NewEchoCustomValidator()
	e.HTTPErrorHandler = web.EchoCustomHTTPErrorHandler
	e.Use(middleware.Logger())

	// DB接続
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:password@localhost/postgres?sslmode=disable"
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}

	// リポジトリとハンドラーの設定
	repo := database.NewFilmSQLRepository(db)
	web.SetupFilmEchoHandlers(e, repo)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// サーバー起動
	e.Logger.Fatal(e.Start(":1323"))
}
