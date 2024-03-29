package Handlers

import (
	"github.com/labstack/echo/v4"
	Requests2 "go-echo-rest-api-sample/src/Presentations/Requests/SuggestedTask"
	"go-echo-rest-api-sample/src/UseCases/SuggestedTask"
	"net/http"
)

type SuggestedTaskHandler struct {
	getSuggestedTasksUseCase UseCases.GetSuggestedTasksUseCase
}

func NewSuggestedTaskHandler(
	getSuggestedTasksUseCase *UseCases.GetSuggestedTasksUseCase,
) *SuggestedTaskHandler {
	return &SuggestedTaskHandler{
		getSuggestedTasksUseCase: *getSuggestedTasksUseCase,
	}
}

func (s *SuggestedTaskHandler) GetSuggestedTasks(ctx echo.Context) error {
	var getSuggestedTasksRequest Requests2.GetSuggestedTasksRequest
	if err := ctx.Bind(&getSuggestedTasksRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	suggestedTasksDto, err := s.getSuggestedTasksUseCase.Execute(getSuggestedTasksRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, suggestedTasksDto)
}
