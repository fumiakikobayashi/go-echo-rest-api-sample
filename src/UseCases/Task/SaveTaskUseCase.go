package UseCases

import (
	"fmt"
	Domains "go-ddd-rest-api-sample/src/Domains/Task"
	"go-ddd-rest-api-sample/src/Presentations/Requests"
	"time"
)

type SaveTaskUseCase struct {
	taskRepository TaskRepositoryInterface
}

func NewSaveTaskUseCase(taskRepository TaskRepositoryInterface) *SaveTaskUseCase {
	return &SaveTaskUseCase{
		taskRepository: taskRepository,
	}
}

func (u *SaveTaskUseCase) Execute(request Requests.SaveTaskRequest) error {
	t, err := time.Parse(Domains.DeadlineFormat, request.Deadline)
	if err != nil {
		return fmt.Errorf("締切日のフォーマットが不正です")
	}

	task := Domains.CreateNewTask(request.Name, t)
	if err := u.taskRepository.SaveTask(task); err != nil {
		return fmt.Errorf("タスクの取得に失敗しました")
	}

	return nil
}
