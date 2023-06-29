package Domains

import (
	"go-ddd-rest-api-sample/src/Infrastructures/Models"
	"go-ddd-rest-api-sample/src/Shared/Errors"
)

func CreateTask(taskModel Models.TaskModel) (*Task, error) {
	id, err := NewTaskId(taskModel.ID)
	if err != nil {
		return &Task{}, Errors.New("001-001", "タスクIDの生成に失敗しました")
	}

	return ReconstructTask(
		id,
		taskModel.Name,
		taskModel.Deadline,
		*taskModel.IsFavorite,
		*taskModel.IsCompleted,
	), nil
}
