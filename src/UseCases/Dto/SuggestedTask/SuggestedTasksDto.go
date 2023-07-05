package Dto

type SuggestedTasksDto struct {
	SuggestedTasksDto []SuggestedTaskDto `json:"suggestedTasks"`
}

func NewSuggestedTasksDto(suggestedTaskDtoList []SuggestedTaskDto) SuggestedTasksDto {
	return SuggestedTasksDto{
		SuggestedTasksDto: suggestedTaskDtoList,
	}
}
