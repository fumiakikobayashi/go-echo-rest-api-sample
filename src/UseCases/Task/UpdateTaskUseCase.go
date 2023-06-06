package UseCases

import (
	"fmt"
	Domains "go-ddd-rest-api-sample/src/Domains/Task"
	"go-ddd-rest-api-sample/src/Presentations/Requests"
	"time"
)

type UpdateTaskUseCase struct {
	taskRepository TaskRepositoryInterface
}

func NewUpdateTaskUseCase(taskRepository TaskRepositoryInterface) *UpdateTaskUseCase {
	return &UpdateTaskUseCase{
		taskRepository: taskRepository,
	}
}

func (u *UpdateTaskUseCase) Execute(request Requests.UpdateTaskRequest) error {
	taskId, _ := Domains.NewTaskId(request.TaskId)
	task, err := u.taskRepository.GetTask(taskId)
	if err != nil {
		return fmt.Errorf("タスクの取得に失敗しました")
	}

	t, err := time.Parse(Domains.DeadlineFormat, request.Deadline)
	if err != nil {
		return fmt.Errorf("締切日のフォーマットが不正です")
	}
	task.UpdateTask(request.Name, t)

	if err := u.taskRepository.UpdateTask(task); err != nil {
		return fmt.Errorf("タスクの更新に失敗しました")
	}
	return nil
}
