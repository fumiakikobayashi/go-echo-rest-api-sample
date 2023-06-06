package Dto

type TaskDto struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Deadline   string `json:"deadline"`
	IsFavorite bool   `json:"isFavorite"`
	IsComplete bool   `json:"isComplete"`
}

func NewTaskDto(
	id int,
	name string,
	deadline string,
	isFavorite bool,
	isComplete bool,
) TaskDto {
	return TaskDto{
		Id:         id,
		Name:       name,
		Deadline:   deadline,
		IsFavorite: isFavorite,
		IsComplete: isComplete,
	}
}
