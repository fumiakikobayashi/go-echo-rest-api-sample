package src

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-ddd-rest-api-sample/src/Infrastructures"
	"go-ddd-rest-api-sample/src/Shared"
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
			panic("DB接続の初期化に失敗しました")
		}
	}(db)

	// ログ設定
	logger := Shared.NewLogger()

	// 依存性の注入したハンドラーを取得
	handlers := NewHandlers(db, logger)

	// echoの初期化
	e := echo.New()

	// カスタムエラーハンドラー
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			_ = c.JSON(he.Code, he.Message)
		}
		//if uce, ok := err.(UseCaseError); ok {
		//	fmt.Println("Caught a use case error:", uce.Message)
		//}
		//else {
		//	_ = c.JSON(http.StatusInternalServerError, err.Error())
		//}
	}

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
