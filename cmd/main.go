package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sashabaranov/go-openai"
	"go-echo-rest-api-sample/src/Infrastructures"
	"go-echo-rest-api-sample/src/Infrastructures/Middlewares"
	"go-echo-rest-api-sample/src/Presentations"
	"go-echo-rest-api-sample/src/Presentations/Routes"
	"go-echo-rest-api-sample/src/Shared/Logger"
	"os"
)

func main() {
	// loggerの初期化
	log := Logger.NewLogger()
	if log == nil {
		log.Error("loggerの初期化に失敗しました")
		return
	}
	defer func(log Logger.ILogger) {
		if err := log.Close(); err != nil {
			log.Error("logファイルのcloseに失敗しました")
			return
		}
	}(log)

	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Error(".envファイルの読み込みに失敗しました")
		return
	}

	// DB接続
	db := Infrastructures.SetupDB()
	if db == nil {
		log.Error("DB接続に失敗しました")
		return
	}
	defer func(db *gorm.DB) {
		if err := db.Close(); err != nil {
			log.Error("DB切断に失敗しました")
			return
		}
	}(db)

	// OpenAIクライアントを生成
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	if client == nil {
		log.Error("OpenAIクライアントの生成に失敗しました")
	}

	// echoの初期化
	e := echo.New()
	if e == nil {
		log.Error("echoの初期化に失敗しました")
		return
	}

	// 依存性の注入したハンドラーを取得
	handlers := Presentations.NewHandlers(db, client)
	if handlers == nil {
		log.Error("ハンドラーの初期化に失敗しました")
		return
	}

	// エンドポイントを定義
	Routes.SetUpRoutes(e, handlers)

	// ミドルウェアの登録
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(Middlewares.CustomLogger(log))

	// Start server
	if err := e.Start(":8080"); err != nil {
		log.Error("サーバーの起動に失敗しました", "err", err.Error())
		return
	}
}
