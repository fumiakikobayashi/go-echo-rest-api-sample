package Domains

import "fmt"

type TaskId struct {
	id int
}

const MinId = 1

func NewTaskId(id int) (TaskId, error) {
	if id < MinId {
		return TaskId{}, fmt.Errorf("TaskIdは1以上の値を入力してください")
	}
	return TaskId{
		id: id,
	}, nil
}

func (t TaskId) GetValue() int {
	return t.id
}
