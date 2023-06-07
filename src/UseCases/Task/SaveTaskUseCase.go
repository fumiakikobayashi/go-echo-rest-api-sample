package UseCases

import (
	"fmt"
	"go-ddd-rest-api-sample/sdk"
	Domains "go-ddd-rest-api-sample/src/Domains/Task"
	Requests "go-ddd-rest-api-sample/src/Presentations/Requests/Task"
	"time"
)

type SaveTaskUseCase struct {
	taskRepository TaskRepositoryInterface
	logger         sdk.LoggerInterface
}

func NewSaveTaskUseCase(taskRepository TaskRepositoryInterface, logger sdk.LoggerInterface) *SaveTaskUseCase {
	return &SaveTaskUseCase{
		taskRepository: taskRepository,
		logger:         logger,
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
