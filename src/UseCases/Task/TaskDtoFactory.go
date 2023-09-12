package UseCases

import (
	Domains "go-echo-rest-api-sample/src/Domains/Task"
	Dto "go-echo-rest-api-sample/src/UseCases/Dto/Task"
)

func CreateTaskDto(task *Domains.Task) Dto.TaskDto {
	return Dto.NewTaskDto(
		task.GetTaskId().GetValue(),
		task.GetName(),
		task.GetDeadline().Format(Domains.DeadlineFormat),
		task.GetIsFavorite(),
		task.GetIsCompleted(),
	)
}
