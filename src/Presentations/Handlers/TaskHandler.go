package Handlers

import (
	"github.com/labstack/echo/v4"
	UseCase "go-ddd-rest-api-sample/src/UseCases/Task"
	"net/http"
)

type TaskHandler struct {
	getTaskUseCase UseCase.GetTaskUseCase
}

func NewTaskHandler(taskUseCase *UseCase.GetTaskUseCase) *TaskHandler {
	return &TaskHandler{
		getTaskUseCase: *taskUseCase,
	}
}

func (c *TaskHandler) GetTasks(ctx echo.Context) error {
	taskListDto, err := c.getTaskUseCase.GetTasks()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, taskListDto)
}
