package Dto

type TaskDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Deadline    string `json:"deadline"`
	IsFavorite  bool   `json:"isFavorite"`
	IsCompleted bool   `json:"isCompleted"`
}

func NewTaskDto(
	id int,
	name string,
	deadline string,
	isFavorite bool,
	isCompleted bool,
) TaskDto {
	return TaskDto{
		Id:          id,
		Name:        name,
		Deadline:    deadline,
		IsFavorite:  isFavorite,
		IsCompleted: isCompleted,
	}
}
