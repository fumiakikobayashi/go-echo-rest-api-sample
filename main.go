package main

import (
	"go-ddd-rest-api-sample/src/Infrastructures"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// DBのセットアップ
	db := Infrastructures.Init()
	defer db.Close()

	// echoのセットアップ
	e := echo.New()

	// ルーティング
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hefrerefefefefg!!!!")
	})

	// サーバーの立ち上げ
	e.Logger.Fatal(e.Start(":8080"))
}
