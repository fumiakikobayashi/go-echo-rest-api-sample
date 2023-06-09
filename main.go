package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-ddd-rest-api-sample/sdk"
	"go-ddd-rest-api-sample/src/Infrastructures"
	"net/http"
)

func main() {
	// .envファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err.Error())
	}

	// DB接続
	db := Infrastructures.Init()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			panic("failed to disconnect database")
		}
	}(db)

	// ログ設定
	logger := sdk.NewLogger()

	// 依存性の注入したハンドラーを取得
	handlers := NewHandlers(db, logger)

	// echoの初期化
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/tasks", handlers.TaskHandler.GetTasks)
	e.GET("/tasks/:taskId", handlers.TaskHandler.GetTask)
	e.POST("/tasks", handlers.TaskHandler.SaveTask)
	e.PUT("/tasks/:taskId", handlers.TaskHandler.UpdateTask)
	e.DELETE("/tasks/:taskId", handlers.TaskHandler.DeleteTask)
	e.PATCH("/tasks/:taskId/favorite", handlers.TaskHandler.UpdateTaskFavorite)
	e.PATCH("/tasks/:taskId/complete", handlers.TaskHandler.UpdateTaskComplete)

	// エラーハンドラー
	e.HTTPErrorHandler = customHTTPErrorHandler

	// Start server
	if err := e.Start(":8080"); err != nil {
		panic(err.Error())
	}
}

func customHTTPErrorHandler(err error, c echo.Context) {
	var code int
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	errorPage := echo.Map{}

	if code == http.StatusNotFound {
		errorPage["message"] = "該当するエンドポイントは見つかりませんでした。"
	}

	if code == http.StatusInternalServerError {
		errorPage["message"] = "予期せぬエラーが発生しました。"
	}

	_ = c.JSON(code, errorPage)
}
