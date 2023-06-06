package sdk

import (
	"github.com/jinzhu/gorm"
	"go-ddd-rest-api-sample/src/Infrastructures/Repositories"
	Presentations "go-ddd-rest-api-sample/src/Presentations/Handlers"
	UseCases "go-ddd-rest-api-sample/src/UseCases/Task"
)

type Handlers struct {
	TaskHandler Presentations.TaskHandler
}

func NewHandlers(db *gorm.DB) *Handlers {
	return &Handlers{
		TaskHandler: *injectTaskHandlerDependencies(db),
	}
}

func injectTaskHandlerDependencies(db *gorm.DB) *Presentations.TaskHandler {
	taskRepository := Repositories.NewTaskRepository(db)
	getTasksUseCase := UseCases.NewGetTasksUseCase(taskRepository)
	getTaskUseCase := UseCases.NewGetTaskUseCase(taskRepository)
	saveTaskUseCase := UseCases.NewSaveTaskUseCase(taskRepository)
	updateTaskUseCase := UseCases.NewUpdateTaskUseCase(taskRepository)
	deleteTaskUseCase := UseCases.NewDeleteTaskUseCase(taskRepository)
	favoriteTaskUseCase := UseCases.NewFavoriteTaskUseCase(taskRepository)
	completeTaskUseCase := UseCases.NewUpdateTaskCompleteUseCase(taskRepository)
	return Presentations.NewTaskHandler(
		getTasksUseCase,
		getTaskUseCase,
		saveTaskUseCase,
		updateTaskUseCase,
		deleteTaskUseCase,
		favoriteTaskUseCase,
		completeTaskUseCase,
	)
}
