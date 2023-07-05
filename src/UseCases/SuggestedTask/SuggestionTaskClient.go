package UseCases

import "go-ddd-rest-api-sample/src/Domains/SuggestedTask"

type SuggestionTaskClient interface {
	SuggestTasksBy(target string) (Domain.SuggestedTasks, error)
}
