package Domains

import (
	"go-echo-rest-api-sample/src/Infrastructures/Models"
	"go-echo-rest-api-sample/src/Shared"
)

func CreateTask(taskModel Models.TaskModel) (*Task, error) {
	id, err := NewTaskId(taskModel.ID)
	if err != nil {
		return &Task{}, Shared.NewSampleError("001-001", "タスクIDの生成に失敗しました")
	}

	return ReconstructTask(
		id,
		taskModel.Name,
		taskModel.Deadline,
		*taskModel.IsFavorite,
		*taskModel.IsCompleted,
	), nil
}
