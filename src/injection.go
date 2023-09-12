package src

import (
	"github.com/jinzhu/gorm"
	"github.com/sashabaranov/go-openai"
	Infrastructures "go-echo-rest-api-sample/src/Infrastructures/Clients"
	"go-echo-rest-api-sample/src/Infrastructures/Repositories"
	Presentations "go-echo-rest-api-sample/src/Presentations/Handlers"
	"go-echo-rest-api-sample/src/Shared"
	UseCases2 "go-echo-rest-api-sample/src/UseCases/SuggestedTask"
	UseCases "go-echo-rest-api-sample/src/UseCases/Task"
)

type Handlers struct {
	TaskHandler          Presentations.TaskHandler
	SuggestedTaskHandler Presentations.SuggestedTaskHandler
}

func NewHandlers(db *gorm.DB, logger *Shared.LoggerInterface, client *openai.Client) *Handlers {
	return &Handlers{
		TaskHandler:          *injectTaskHandlerDependencies(db, logger),
		SuggestedTaskHandler: *injectSuggestedTaskHandlerDependencies(client, logger),
	}
}

func injectTaskHandlerDependencies(db *gorm.DB, logger *Shared.LoggerInterface) *Presentations.TaskHandler {
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

func injectSuggestedTaskHandlerDependencies(client *openai.Client, logger *Shared.LoggerInterface) *Presentations.SuggestedTaskHandler {
	suggestionTaskClient := Infrastructures.NewSuggestionTaskClient(logger, client)
	GetSuggestedTaskUseCase := UseCases2.NewGetSuggestedTasksUseCase(suggestionTaskClient)
	return Presentations.NewSuggestedTaskHandler(
		GetSuggestedTaskUseCase,
		logger,
	)
}
