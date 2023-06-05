package Domains

type TaskId struct {
	id int
}

func NewTaskId(id int) (TaskId, error) {
	return TaskId{
		id: id,
	}, nil
}

func (t TaskId) GetValue() int {
	return t.id
}
