package Handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go-echo-rest-api-sample/src/Presentations/Requests/Task"
	"go-echo-rest-api-sample/src/Shared/Logger"
	"go-echo-rest-api-sample/src/UseCases/Task"
	"net/http"
)

type TaskHandler struct {
	getTasksUseCase           UseCases.GetTasksUseCase
	getTaskUseCase            UseCases.GetTaskUseCase
	saveTaskUseCase           UseCases.SaveTaskUseCase
	updateTaskUseCase         UseCases.UpdateTaskUseCase
	deleteTaskUseCase         UseCases.DeleteTaskUseCase
	updateFavoriteTaskUseCase UseCases.UpdateTaskFavoriteUseCase
	updateCompleteTaskUseCase UseCases.UpdateTaskCompleteUseCase
}

func NewTaskHandler(
	getTasksUseCase *UseCases.GetTasksUseCase,
	getTaskUseCase *UseCases.GetTaskUseCase,
	saveTaskUseCase *UseCases.SaveTaskUseCase,
	updateTaskUseCase *UseCases.UpdateTaskUseCase,
	deleteTaskUseCase *UseCases.DeleteTaskUseCase,
	updateFavoriteTaskUseCase *UseCases.UpdateTaskFavoriteUseCase,
	updateCompleteTaskUseCase *UseCases.UpdateTaskCompleteUseCase,
) *TaskHandler {
	return &TaskHandler{
		getTasksUseCase:           *getTasksUseCase,
		getTaskUseCase:            *getTaskUseCase,
		saveTaskUseCase:           *saveTaskUseCase,
		updateTaskUseCase:         *updateTaskUseCase,
		deleteTaskUseCase:         *deleteTaskUseCase,
		updateFavoriteTaskUseCase: *updateFavoriteTaskUseCase,
		updateCompleteTaskUseCase: *updateCompleteTaskUseCase,
	}
}

func (c *TaskHandler) GetTasks(ctx echo.Context) error {
	logger := Logger.FromContext(ctx)

	var tasksRequest Requests.GetTasksRequest
	if err := ctx.Bind(&tasksRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "リクエストパラメータのバインドに失敗しました")
	}

	if err := tasksRequest.Validate(); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			logger.Error("バリデーションに失敗しました", "Field", err.Field(), "Tag", err.Tag(), "ActualValue", err.Value(), "ExpectedValue", err.Param())
		}
		return echo.NewHTTPError(http.StatusBadRequest, "リクエストパラメータが不正です")
	}

	taskListDto, err := c.getTasksUseCase.Execute(tasksRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "タスク一覧の取得に失敗しました")
	}

	return ctx.JSON(http.StatusOK, taskListDto)
}

func (c *TaskHandler) GetTask(ctx echo.Context) error {
	logger := Logger.FromContext(ctx)

	var taskRequest Requests.GetTaskRequest
	if err := ctx.Bind(&taskRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := taskRequest.Validate(); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			logger.Error("バリデーションに失敗しました", "Field", err.Field(), "Tag", err.Tag(), "ActualValue", err.Value(), "ExpectedValue", err.Param())
		}
		return echo.NewHTTPError(http.StatusBadRequest, "リクエストパラメータが不正です")
	}

	taskDto, err := c.getTaskUseCase.Execute(taskRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, taskDto)
}

func (c *TaskHandler) SaveTask(ctx echo.Context) error {
	logger := Logger.FromContext(ctx)

	var taskRequest Requests.SaveTaskRequest
	if err := ctx.Bind(&taskRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := taskRequest.Validate(); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			logger.Error("バリデーションに失敗しました", "Field", err.Field(), "Tag", err.Tag(), "ActualValue", err.Value(), "ExpectedValue", err.Param())
		}
		return echo.NewHTTPError(http.StatusBadRequest, "リクエストパラメータが不正です")
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
	logger := Logger.FromContext(ctx)

	var taskRequest Requests.UpdateTaskRequest
	if err := ctx.Bind(&taskRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := taskRequest.Validate(); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			logger.Error("バリデーションに失敗しました", "Field", err.Field(), "Tag", err.Tag(), "ActualValue", err.Value(), "ExpectedValue", err.Param())
		}
		return echo.NewHTTPError(http.StatusBadRequest, "リクエストパラメータが不正です")
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
	logger := Logger.FromContext(ctx)

	var taskRequest Requests.DeleteTaskRequest
	if err := ctx.Bind(&taskRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.deleteTaskUseCase.Execute(taskRequest); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := taskRequest.Validate(); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			logger.Error("バリデーションに失敗しました", "Field", err.Field(), "Tag", err.Tag(), "ActualValue", err.Value(), "ExpectedValue", err.Param())
		}
		return echo.NewHTTPError(http.StatusBadRequest, "リクエストパラメータが不正です")
	}

	response := echo.Map{
		"message": "タスクを削除しました",
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *TaskHandler) UpdateTaskFavorite(ctx echo.Context) error {
	logger := Logger.FromContext(ctx)

	var taskRequest Requests.UpdateTaskFavoriteRequest
	if err := ctx.Bind(&taskRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.updateFavoriteTaskUseCase.Execute(taskRequest); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := taskRequest.Validate(); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			logger.Error("バリデーションに失敗しました", "Field", err.Field(), "Tag", err.Tag(), "ActualValue", err.Value(), "ExpectedValue", err.Param())
		}
		return echo.NewHTTPError(http.StatusBadRequest, "リクエストパラメータが不正です")
	}

	response := echo.Map{
		"message": "お気に入り状態を更新しました",
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *TaskHandler) UpdateTaskComplete(ctx echo.Context) error {
	logger := Logger.FromContext(ctx)

	var taskRequest Requests.UpdateTaskCompleteRequest
	if err := ctx.Bind(&taskRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := taskRequest.Validate(); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			logger.Error("バリデーションに失敗しました", "Field", err.Field(), "Tag", err.Tag(), "ActualValue", err.Value(), "ExpectedValue", err.Param())
		}
		return echo.NewHTTPError(http.StatusBadRequest, "リクエストパラメータが不正です")
	}

	if err := c.updateCompleteTaskUseCase.Execute(taskRequest); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := echo.Map{
		"message": "完了状態を更新しました",
	}
	return ctx.JSON(http.StatusOK, response)
}
