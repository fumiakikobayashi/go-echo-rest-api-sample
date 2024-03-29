package Domains

import (
	"go-echo-rest-api-sample/src/Shared"
)

type TaskList struct {
	taskList map[TaskId]*Task
}

func NewTaskList() *TaskList {
	return &TaskList{
		taskList: make(map[TaskId]*Task),
	}
}

func (t *TaskList) Push(task *Task) error {
	if _, ok := t.taskList[task.GetTaskId()]; ok {
		return Shared.NewSampleError("001-001", "すでに存在しているtaskIdをpushしようとしています。")
	}
	t.taskList[task.GetTaskId()] = task
	return nil
}

func (t *TaskList) GetTaskList() map[TaskId]*Task {
	return t.taskList
}
