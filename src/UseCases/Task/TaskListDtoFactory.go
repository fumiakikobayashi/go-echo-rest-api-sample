package UseCases

import (
	Domains "go-ddd-rest-api-sample/src/Domains/Task"
	Dto "go-ddd-rest-api-sample/src/UseCases/Dto/Task"
)

func CreateTaskDtoList(taskList *Domains.TaskList) (Dto.TaskListDto, error) {
	var dtoList []Dto.TaskDto
	for _, task := range taskList.GetTaskList() {
		dto, err := CreateTaskDto(task)
		if err != nil {
			return Dto.TaskListDto{}, err
		}
		dtoList = append(dtoList, dto)
	}

	return Dto.NewTaskDtoList(dtoList), nil
}
