package Domains

import (
	"go-ddd-rest-api-sample/src/Shared/Errors"
)

type TaskId struct {
	id int
}

const MIN = 1

func NewTaskId(id int) (TaskId, error) {
	if id < MIN {
		return TaskId{}, Errors.New("001-001", "タスクIDは0以上の値を指定してください")
	}
	return TaskId{
		id: id,
	}, nil
}

func (t TaskId) GetValue() int {
	return t.id
}
