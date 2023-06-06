package Handlers

import (
	"github.com/labstack/echo/v4"
	"go-ddd-rest-api-sample/src/Presentations/Requests"
	UseCase "go-ddd-rest-api-sample/src/UseCases/Task"
	"net/http"
)

type TaskHandler struct {
	getTasksUseCase   UseCase.GetTasksUseCase
	getTaskUseCase    UseCase.GetTaskUseCase
	saveTaskUseCase   UseCase.SaveTaskUseCase
	updateTaskUseCase UseCase.UpdateTaskUseCase
	deleteTaskUseCase UseCase.DeleteTaskUseCase
}

func NewTaskHandler(
	getTasksUseCase *UseCase.GetTasksUseCase,
	getTaskUseCase *UseCase.GetTaskUseCase,
	saveTaskUseCase *UseCase.SaveTaskUseCase,
	updateTaskUseCase *UseCase.UpdateTaskUseCase,
	deleteTaskUseCase *UseCase.DeleteTaskUseCase,
) *TaskHandler {
	return &TaskHandler{
		getTasksUseCase:   *getTasksUseCase,
		getTaskUseCase:    *getTaskUseCase,
		saveTaskUseCase:   *saveTaskUseCase,
		updateTaskUseCase: *updateTaskUseCase,
		deleteTaskUseCase: *deleteTaskUseCase,
	}
}

func (c *TaskHandler) GetTasks(ctx echo.Context) error {
	taskListDto, err := c.getTasksUseCase.Execute()
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

	taskDto, err := c.getTaskUseCase.Execute(taskRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, taskDto)
}

func (c *TaskHandler) SaveTask(ctx echo.Context) error {
	var taskRequest Requests.SaveTaskRequest
	if err := ctx.Bind(&taskRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.saveTaskUseCase.Execute(taskRequest); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := echo.Map{
		"message": "タスクを保存しました",
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *TaskHandler) UpdateTask(ctx echo.Context) error {
	var taskRequest Requests.UpdateTaskRequest
	if err := ctx.Bind(&taskRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.updateTaskUseCase.Execute(taskRequest); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := echo.Map{
		"message": "タスクを更新しました",
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *TaskHandler) DeleteTask(ctx echo.Context) error {
	var taskRequest Requests.DeleteTaskRequest
	if err := ctx.Bind(&taskRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.deleteTaskUseCase.Execute(taskRequest); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := echo.Map{
		"message": "タスクを削除しました",
	}
	return ctx.JSON(http.StatusOK, response)
}
