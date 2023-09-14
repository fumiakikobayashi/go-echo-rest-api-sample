package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sashabaranov/go-openai"
	"go-echo-rest-api-sample/src/Infrastructures"
	"go-echo-rest-api-sample/src/Presentations"
	"go-echo-rest-api-sample/src/Presentations/Routes"
	"os"
)

func main() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		panic("Envファイルの読み込みに失敗しました")
	}

	// DB接続
	db := Infrastructures.Init()
	defer func(db *gorm.DB) {
		if err := db.Close(); err != nil {
			panic("DB接続の初期化に失敗しました")
		}
	}(db)

	// OpenAIのクライアント
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	// echoの初期化
	e := echo.New()

	// 依存性の注入したハンドラーを取得
	handlers := Presentations.NewHandlers(db, client)

	// エンドポイントを定義
	Routes.SetUpRoutes(e, handlers)

	// ミドルウェアの登録
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Start server
	if err := e.Start(":8080"); err != nil {
		panic(err.Error())
	}
}
