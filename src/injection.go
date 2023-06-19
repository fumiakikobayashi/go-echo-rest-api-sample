package src

import (
	"github.com/jinzhu/gorm"
	"go-ddd-rest-api-sample/src/Infrastructures/Repositories"
	Presentations "go-ddd-rest-api-sample/src/Presentations/Handlers"
	"go-ddd-rest-api-sample/src/Shared"
	UseCases "go-ddd-rest-api-sample/src/UseCases/Task"
)

type Handlers struct {
	TaskHandler Presentations.TaskHandler
}

func NewHandlers(db *gorm.DB, logger Shared.LoggerInterface) *Handlers {
	return &Handlers{
		TaskHandler: *injectTaskHandlerDependencies(db, logger),
	}
}

func injectTaskHandlerDependencies(db *gorm.DB, logger Shared.LoggerInterface) *Presentations.TaskHandler {
	taskRepository := Repositories.NewTaskRepository(db, logger)
	getTasksUseCase := UseCases.NewGetTasksUseCase(taskRepository, logger)
	getTaskUseCase := UseCases.NewGetTaskUseCase(taskRepository, logger)
	saveTaskUseCase := UseCases.NewSaveTaskUseCase(taskRepository, logger)
	updateTaskUseCase := UseCases.NewUpdateTaskUseCase(taskRepository, logger)
	deleteTaskUseCase := UseCases.NewDeleteTaskUseCase(taskRepository, logger)
	favoriteTaskUseCase := UseCases.NewFavoriteTaskUseCase(taskRepository, logger)
	completeTaskUseCase := UseCases.NewUpdateTaskCompleteUseCase(taskRepository, logger)
	return Presentations.NewTaskHandler(
		getTasksUseCase,
		getTaskUseCase,
		saveTaskUseCase,
		updateTaskUseCase,
		deleteTaskUseCase,
		favoriteTaskUseCase,
		completeTaskUseCase,
		logger,
	)
}
