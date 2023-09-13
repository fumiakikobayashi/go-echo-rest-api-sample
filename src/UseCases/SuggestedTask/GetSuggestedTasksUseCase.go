package UseCases

import (
	"go-echo-rest-api-sample/src/Presentations/Requests/SuggestedTask"
	"go-echo-rest-api-sample/src/UseCases/Dto/SuggestedTask"
)

type GetSuggestedTasksUseCase struct {
	suggestionTask SuggestionTaskClient
}

func NewGetSuggestedTasksUseCase(suggestionTask SuggestionTaskClient) *GetSuggestedTasksUseCase {
	return &GetSuggestedTasksUseCase{
		suggestionTask: suggestionTask,
	}
}

func (uc *GetSuggestedTasksUseCase) Execute(request Requests.GetSuggestedTasksRequest) (Dto.SuggestedTasksDto, error) {
	suggestedTasks, err := uc.suggestionTask.SuggestTasksBy(request.Target)
	if err != nil {
		return Dto.SuggestedTasksDto{}, err
	}
	return CreateSuggestedTasksDto(suggestedTasks), nil
}
