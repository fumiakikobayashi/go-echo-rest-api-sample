package UseCases

import (
	Domains "go-ddd-rest-api-sample/src/Domains/Task"
	Requests "go-ddd-rest-api-sample/src/Presentations/Requests/Task"
	"go-ddd-rest-api-sample/src/Shared"
	"time"
)

type SaveTaskUseCase struct {
	taskRepository TaskRepositoryInterface
	logger         Shared.LoggerInterface
}

func NewSaveTaskUseCase(taskRepository TaskRepositoryInterface, logger Shared.LoggerInterface) *SaveTaskUseCase {
	return &SaveTaskUseCase{
		taskRepository: taskRepository,
		logger:         logger,
	}
}

func (u *SaveTaskUseCase) Execute(request Requests.SaveTaskRequest) error {
	t, err := time.Parse(Domains.DeadlineFormat, request.Deadline)
	if err != nil {
		return Shared.New("001-001", "締切日のフォーマットが不正です")
	}

	task := Domains.CreateNewTask(request.Name, t)
	if err := u.taskRepository.SaveTask(task); err != nil {
		return err
	}

	return nil
}
