package Domains

import (
	"go-ddd-rest-api-sample/src/Shared"
)

type TaskId struct {
	id int
}

const MIN = 1

func NewTaskId(id int) (TaskId, error) {
	if id < MIN {
		return TaskId{}, Shared.NewSampleError("001-001", "タスクIDは0以上の値を指定してください")
	}
	return TaskId{
		id: id,
	}, nil
}

func (t TaskId) GetValue() int {
	return t.id
}
