package Routes

import (
	"github.com/labstack/echo/v4"
	"go-echo-rest-api-sample/src/Presentations"
)

func SetUpRoutes(e *echo.Echo, handlers *Presentations.Handlers) {
	e.GET("/tasks", handlers.TaskHandler.GetTasks)
	e.GET("/tasks/:taskId", handlers.TaskHandler.GetTask)
	e.POST("/tasks", handlers.TaskHandler.SaveTask)
	e.PUT("/tasks/:taskId", handlers.TaskHandler.UpdateTask)
	e.DELETE("/tasks/:taskId", handlers.TaskHandler.DeleteTask)
	e.PATCH("/tasks/:taskId/favorite", handlers.TaskHandler.UpdateTaskFavorite)
	e.PATCH("/tasks/:taskId/complete", handlers.TaskHandler.UpdateTaskComplete)
	e.GET("/tasks/suggestion", handlers.SuggestedTaskHandler.GetSuggestedTasks)
}
