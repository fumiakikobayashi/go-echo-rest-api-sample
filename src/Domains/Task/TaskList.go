package Domains

import "fmt"

type TaskList struct {
	taskList map[TaskId]*Task
}

func NewTaskList() TaskList {
	return TaskList{
		taskList: make(map[TaskId]*Task),
	}
}

func (t *TaskList) Push(task *Task) error {
	if _, ok := t.taskList[task.GetTaskId()]; ok {
		return fmt.Errorf("タスクIDが重複しています。")
	}
	t.taskList[task.GetTaskId()] = task
	return nil
}

func (t *TaskList) GetTaskList() map[TaskId]*Task {
	return t.taskList
}
