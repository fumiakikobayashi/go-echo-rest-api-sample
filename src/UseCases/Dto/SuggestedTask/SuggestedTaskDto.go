package Dto

type SuggestedTaskDto struct {
	Name string `json:"name"`
}

func NewSuggestedTaskDto(name string) SuggestedTaskDto {
	return SuggestedTaskDto{
		Name: name,
	}
}
