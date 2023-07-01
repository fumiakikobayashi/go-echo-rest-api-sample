package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-ddd-rest-api-sample/src"
	"go-ddd-rest-api-sample/src/Infrastructures"
	"go-ddd-rest-api-sample/src/Shared"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	// .envファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Errors loading .env file")
		panic(err.Error())
	}

	// DB接続
	db := Infrastructures.Init()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			panic("DB接続の初期化に失敗しました")
		}
	}(db)

	// ログ設定
	logger := Shared.NewLogger()

	// 依存性の注入したハンドラーを取得
	handlers := src.NewHandlers(db, logger)

	// echoの初期化
	e := echo.New()

	// カスタムエラーハンドラー
	e.HTTPErrorHandler = customHTTPErrorHandler

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

	// Start server
	if err := e.Start(":8080"); err != nil {
		panic(err.Error())
	}
}

func customHTTPErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if !ok {
		zap.S().Errorf("Unknown error: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	httpCode := he.Code
	switch err := he.Message.(type) {
	case error:
		switch {
		case httpCode >= 500:
			zap.S().Errorf("Server error: %v", err)
			if me, ok := err.(*Shared.SampleError); ok {
				fmt.Print(me.StackTrace)
			}
		case httpCode >= 400:
			zap.S().Infof("Client error: %v", err)
		}
		c.JSON(httpCode, "error")
	case string:
		// 存在しないエンドポイントが叩かれた場合
		zap.S().Errorf("Echo HTTP error: %v", he)
		c.JSON(http.StatusInternalServerError, he)
	default:
		zap.S().Errorf("Unknown HTTP error: %v", he)
		c.JSON(http.StatusInternalServerError, "予期せぬエラーが発生しました")
	}
}
