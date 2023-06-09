package Domains

import "go-ddd-rest-api-sample/src/Infrastructures/Models"

func CreateTask(taskModel Models.TaskModel) (*Task, error) {
	id, err := NewTaskId(taskModel.ID)
	if err != nil {
		return &Task{}, err
	}

	return ReconstructTask(
		id,
		taskModel.Name,
		taskModel.Deadline,
		*taskModel.IsFavorite,
		*taskModel.IsCompleted,
	), nil
}
