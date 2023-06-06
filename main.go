package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"go-ddd-rest-api-sample/src/Infrastructures"
	"go-ddd-rest-api-sample/src/Infrastructures/Repositories"
	"go-ddd-rest-api-sample/src/Presentations/Handlers"
	UseCases "go-ddd-rest-api-sample/src/UseCases/Task"
	"net/http"
	"os"
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
	logLevelMap := map[string]zerolog.Level{
		"debug": zerolog.DebugLevel,
		"info":  zerolog.InfoLevel,
		"warn":  zerolog.WarnLevel,
		"error": zerolog.ErrorLevel,
		"fatal": zerolog.FatalLevel,
		"panic": zerolog.PanicLevel,
	}
	if level, ok := logLevelMap[os.Getenv("LOG_LEVEL")]; ok {
		zerolog.SetGlobalLevel(level)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// DI（依存性の注入）
	taskRepository := Repositories.NewTaskRepository(db)
	getTasksUseCase := UseCases.NewGetTasksUseCase(taskRepository)
	getTaskUseCase := UseCases.NewGetTaskUseCase(taskRepository)
	saveTaskUseCase := UseCases.NewSaveTaskUseCase(taskRepository)
	updateTaskUseCase := UseCases.NewUpdateTaskUseCase(taskRepository)
	deleteTaskUseCase := UseCases.NewDeleteTaskUseCase(taskRepository)
	taskHandler := Handlers.NewTaskHandler(
		getTasksUseCase,
		getTaskUseCase,
		saveTaskUseCase,
		updateTaskUseCase,
		deleteTaskUseCase,
	)

	// echoの初期化
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/tasks", taskHandler.GetTasks)
	e.GET("/tasks/:taskId", taskHandler.GetTask)
	e.POST("/tasks", taskHandler.SaveTask)
	e.PUT("/tasks/:taskId", taskHandler.UpdateTask)
	e.PUT("/tasks/:taskId", taskHandler.DeleteTask)

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
