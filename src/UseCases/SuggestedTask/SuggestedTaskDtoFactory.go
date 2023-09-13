package UseCases

import (
	Domain "go-echo-rest-api-sample/src/Domains/SuggestedTask"
	"go-echo-rest-api-sample/src/UseCases/Dto/SuggestedTask"
)

func CreateSuggestedTaskDto(task Domain.SuggestedTask) Dto.SuggestedTaskDto {
	return Dto.NewSuggestedTaskDto(
		task.GetName(),
	)
}
