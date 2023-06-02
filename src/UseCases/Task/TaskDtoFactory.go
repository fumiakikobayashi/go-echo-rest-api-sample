package UseCases

import (
	Domains "go-ddd-rest-api-sample/src/Domains/Task"
	Dto "go-ddd-rest-api-sample/src/UseCases/Dto/Task"
)

func CreateTaskDto(task *Domains.Task) (Dto.TaskDto, error) {
	return Dto.NewTaskDto(
		task.GetTaskId().GetValue(),
		task.GetName(),
		task.GetDeadline().Format("2006-01-02"),
		task.GetIsFavorite(),
		task.GetIsComplete(),
	), nil
}
