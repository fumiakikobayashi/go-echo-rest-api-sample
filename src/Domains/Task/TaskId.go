package Domains

import "fmt"

type TaskId struct {
	id int
}

const MIN = 1

func NewTaskId(id int) (TaskId, error) {
	if id < MIN {
		return TaskId{}, fmt.Errorf("タスクIDは0以上の値を指定してください")
	}
	return TaskId{
		id: id,
	}, nil
}

func (t TaskId) GetValue() int {
	return t.id
}
