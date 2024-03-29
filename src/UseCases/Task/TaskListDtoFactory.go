package UseCases

import (
	"go-echo-rest-api-sample/src/Domains/Task"
	"go-echo-rest-api-sample/src/UseCases/Dto/Task"
)

func CreateTaskDtoList(taskList *Domains.TaskList) Dto.TaskListDto {
	var dtoList []Dto.TaskDto
	for _, task := range taskList.GetTaskList() {
		dto := CreateTaskDto(task)
		dtoList = append(dtoList, dto)
	}

	return Dto.NewTaskDtoList(dtoList)
}
