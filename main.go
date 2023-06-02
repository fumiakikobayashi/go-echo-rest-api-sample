package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-ddd-rest-api-sample/src/Infrastructures"
	"go-ddd-rest-api-sample/src/Infrastructures/Repositories"
	"go-ddd-rest-api-sample/src/Presentations/Handlers"
	UseCases "go-ddd-rest-api-sample/src/UseCases/Task"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err.Error())
	}

	db := Infrastructures.Init()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			panic("failed to disconnect database")
		}
	}(db)

	// DI（依存性の注入）
	taskRepository := Repositories.NewTaskRepository(db)
	taskUseCase := UseCases.NewGetTaskUseCase(taskRepository)
	taskController := Handlers.NewTaskHandler(taskUseCase)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/tasks", taskController.GetTasks)

	e.Start(":8080")
}
