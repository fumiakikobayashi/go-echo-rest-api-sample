package Dto

type TaskListDto struct {
	TaskDtoList []TaskDto
}

func NewTaskDtoList(taskDtoList []TaskDto) TaskListDto {
	return TaskListDto{
		TaskDtoList: taskDtoList,
	}
}
