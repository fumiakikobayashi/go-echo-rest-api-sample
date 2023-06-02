package Dto

type TaskDto struct {
	Id         int
	Name       string
	Deadline   string
	IsFavorite bool
	IsComplete bool
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
