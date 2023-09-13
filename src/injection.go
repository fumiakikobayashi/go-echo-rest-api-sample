package src

import (
	"github.com/jinzhu/gorm"
	"github.com/sashabaranov/go-openai"
	"go-echo-rest-api-sample/src/Infrastructures/Clients"
	"go-echo-rest-api-sample/src/Infrastructures/Repositories"
	Handlers3 "go-echo-rest-api-sample/src/Presentations/Handlers"
	UseCases2 "go-echo-rest-api-sample/src/UseCases/SuggestedTask"
	"go-echo-rest-api-sample/src/UseCases/Task"
)

type Handlers struct {
	TaskHandler          Handlers3.TaskHandler
	SuggestedTaskHandler Handlers3.SuggestedTaskHandler
}

func NewHandlers(db *gorm.DB, client *openai.Client) *Handlers {
	return &Handlers{
		TaskHandler:          *injectTaskHandlerDependencies(db),
		SuggestedTaskHandler: *injectSuggestedTaskHandlerDependencies(client),
	}
}

func injectTaskHandlerDependencies(db *gorm.DB) *Handlers3.TaskHandler {
	taskRepository := Repositories.NewTaskRepository(db)
	getTasksUseCase := UseCases.NewGetTasksUseCase(taskRepository)
	getTaskUseCase := UseCases.NewGetTaskUseCase(taskRepository)
	saveTaskUseCase := UseCases.NewSaveTaskUseCase(taskRepository)
	updateTaskUseCase := UseCases.NewUpdateTaskUseCase(taskRepository)
	deleteTaskUseCase := UseCases.NewDeleteTaskUseCase(taskRepository)
	favoriteTaskUseCase := UseCases.NewFavoriteTaskUseCase(taskRepository)
	completeTaskUseCase := UseCases.NewUpdateTaskCompleteUseCase(taskRepository)
	return Handlers3.NewTaskHandler(
		getTasksUseCase,
		getTaskUseCase,
		saveTaskUseCase,
		updateTaskUseCase,
		deleteTaskUseCase,
		favoriteTaskUseCase,
		completeTaskUseCase,
	)
}

func injectSuggestedTaskHandlerDependencies(client *openai.Client) *Handlers3.SuggestedTaskHandler {
	suggestionTaskClient := Infrastructures.NewSuggestionTaskClient(client)
	GetSuggestedTaskUseCase := UseCases2.NewGetSuggestedTasksUseCase(suggestionTaskClient)
	return Handlers3.NewSuggestedTaskHandler(GetSuggestedTaskUseCase)
}
