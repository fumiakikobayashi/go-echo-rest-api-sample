package Presentations

import (
	"github.com/jinzhu/gorm"
	"github.com/sashabaranov/go-openai"
	"go-echo-rest-api-sample/src/Infrastructures/Clients"
	"go-echo-rest-api-sample/src/Infrastructures/Repositories"
	Handlers2 "go-echo-rest-api-sample/src/Presentations/Handlers"
	UseCases2 "go-echo-rest-api-sample/src/UseCases/SuggestedTask"
	"go-echo-rest-api-sample/src/UseCases/Task"
)

type Handlers struct {
	TaskHandler          Handlers2.TaskHandler
	SuggestedTaskHandler Handlers2.SuggestedTaskHandler
}

func NewHandlers(db *gorm.DB, client *openai.Client) *Handlers {
	return &Handlers{
		TaskHandler:          *injectTaskHandlerDependencies(db),
		SuggestedTaskHandler: *injectSuggestedTaskHandlerDependencies(client),
	}
}

func injectTaskHandlerDependencies(db *gorm.DB) *Handlers2.TaskHandler {
	taskRepository := Repositories.NewTaskRepository(db)
	getTasksUseCase := UseCases.NewGetTasksUseCase(taskRepository)
	getTaskUseCase := UseCases.NewGetTaskUseCase(taskRepository)
	saveTaskUseCase := UseCases.NewSaveTaskUseCase(taskRepository)
	updateTaskUseCase := UseCases.NewUpdateTaskUseCase(taskRepository)
	deleteTaskUseCase := UseCases.NewDeleteTaskUseCase(taskRepository)
	favoriteTaskUseCase := UseCases.NewFavoriteTaskUseCase(taskRepository)
	completeTaskUseCase := UseCases.NewUpdateTaskCompleteUseCase(taskRepository)
	return Handlers2.NewTaskHandler(
		getTasksUseCase,
		getTaskUseCase,
		saveTaskUseCase,
		updateTaskUseCase,
		deleteTaskUseCase,
		favoriteTaskUseCase,
		completeTaskUseCase,
	)
}

func injectSuggestedTaskHandlerDependencies(client *openai.Client) *Handlers2.SuggestedTaskHandler {
	suggestionTaskClient := Infrastructures.NewSuggestionTaskClient(client)
	GetSuggestedTaskUseCase := UseCases2.NewGetSuggestedTasksUseCase(suggestionTaskClient)
	return Handlers2.NewSuggestedTaskHandler(GetSuggestedTaskUseCase)
}
