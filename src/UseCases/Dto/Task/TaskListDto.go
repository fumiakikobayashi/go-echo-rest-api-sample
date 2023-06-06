package Dto

type TaskListDto struct {
	TaskList []TaskDto `json:"tasks"`
}

func NewTaskDtoList(taskDtoList []TaskDto) TaskListDto {
	return TaskListDto{
		TaskList: taskDtoList,
	}
}
