package Handlers

import (
	"github.com/labstack/echo/v4"
	"go-ddd-rest-api-sample/src/Presentations/Requests"
	UseCase "go-ddd-rest-api-sample/src/UseCases/Task"
	"net/http"
)

type TaskHandler struct {
	getTasksUseCase UseCase.GetTasksUseCase
	getTaskUseCase  UseCase.GetTaskUseCase
}

func NewTaskHandler(
	getTasksUseCase *UseCase.GetTasksUseCase,
	getTaskUseCase *UseCase.GetTaskUseCase,
) *TaskHandler {
	return &TaskHandler{
		getTasksUseCase: *getTasksUseCase,
		getTaskUseCase:  *getTaskUseCase,
	}
}

func (c *TaskHandler) GetTasks(ctx echo.Context) error {
	taskListDto, err := c.getTasksUseCase.GetTasks()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, taskListDto)
}

func (c *TaskHandler) GetTask(ctx echo.Context) error {
	var taskRequest Requests.GetTaskRequest
	if err := ctx.Bind(&taskRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	taskDto, err := c.getTaskUseCase.GetTask(taskRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, taskDto)
}
