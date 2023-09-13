package UseCases

import (
	"go-echo-rest-api-sample/src/Domains/SuggestedTask"
	"go-echo-rest-api-sample/src/UseCases/Dto/SuggestedTask"
)

func CreateSuggestedTasksDto(suggestedTasks Domain.SuggestedTasks) Dto.SuggestedTasksDto {
	var suggestedTaskDtoList []Dto.SuggestedTaskDto
	for _, suggestedTask := range suggestedTasks.GetSuggestedTasks() {
		dto := CreateSuggestedTaskDto(suggestedTask)
		suggestedTaskDtoList = append(suggestedTaskDtoList, dto)
	}
	return Dto.NewSuggestedTasksDto(suggestedTaskDtoList)
}
