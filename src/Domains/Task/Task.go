package Domains

import "time"

type Task struct {
	taskId      TaskId
	name        string
	deadline    time.Time
	isFavorite  bool
	isCompleted bool
}

const DeadlineFormat = "2006-01-02"

func CreateNewTask(name string, deadline time.Time) *Task {
	return &Task{
		name:        name,
		deadline:    deadline,
		isFavorite:  false,
		isCompleted: false,
	}
}

func ReconstructTask(taskId TaskId, name string, deadline time.Time, isFavorite bool, isCompleted bool) *Task {
	return &Task{
		taskId:      taskId,
		name:        name,
		deadline:    deadline,
		isFavorite:  isFavorite,
		isCompleted: isCompleted,
	}
}

func (t *Task) UpdateTask(name string, deadline time.Time) {
	t.name = name
	t.deadline = deadline
}

func (t *Task) UpdateTaskFavorite() {
	t.isFavorite = !t.isFavorite
}

func (t *Task) UpdateTaskComplete() {
	t.isCompleted = !t.isCompleted
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

func (t *Task) GetIsCompleted() bool {
	return t.isCompleted
}
