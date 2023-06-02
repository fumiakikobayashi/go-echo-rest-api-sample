package Domains

import "time"

type Task struct {
	taskId     TaskId
	name       string
	deadline   time.Time
	isFavorite bool
	isComplete bool
}

func NewTask(taskId TaskId, name string, deadline time.Time, isFavorite bool, isComplete bool) *Task {
	return &Task{
		taskId:     taskId,
		name:       name,
		deadline:   deadline,
		isFavorite: isFavorite,
		isComplete: isComplete,
	}
}

func (t *Task) GetTaskId() TaskId {
	return t.taskId
}

func (t *Task) GetName() string {
	return t.name
}

func (t *Task) GetDeadline() time.Time {
	return t.deadline
}

func (t *Task) GetIsFavorite() bool {
	return t.isFavorite
}

func (t *Task) GetIsComplete() bool {
	return t.isComplete
}
