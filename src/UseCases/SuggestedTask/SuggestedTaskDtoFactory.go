package UseCases

import (
	Domain "go-ddd-rest-api-sample/src/Domains/SuggestedTask"
	Dto "go-ddd-rest-api-sample/src/UseCases/Dto/SuggestedTask"
)

func CreateSuggestedTaskDto(task Domain.SuggestedTask) Dto.SuggestedTaskDto {
	return Dto.NewSuggestedTaskDto(
		task.GetName(),
	)
}
