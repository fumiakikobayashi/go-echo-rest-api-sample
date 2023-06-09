package UseCases

import (
	Domains "go-ddd-rest-api-sample/src/Domains/Task"
	Dto "go-ddd-rest-api-sample/src/UseCases/Dto/Task"
)

func CreateTaskDto(task *Domains.Task) (Dto.TaskDto, error) {
	return Dto.NewTaskDto(
		task.GetTaskId().GetValue(),
		task.GetName(),
		task.GetDeadline().Format(Domains.DeadlineFormat),
		task.GetIsFavorite(),
		task.GetIsCompleted(),
	), nil
}
