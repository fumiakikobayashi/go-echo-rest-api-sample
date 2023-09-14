package DomainServices

import Domain "go-echo-rest-api-sample/src/Domains/SuggestedTask"

type ISuggestionTask interface {
	SuggestTasksBy(target string) (Domain.SuggestedTasks, error)
}
