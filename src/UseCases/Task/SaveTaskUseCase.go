package UseCases

import (
	"go-echo-rest-api-sample/src/DomainServices"
	"go-echo-rest-api-sample/src/Domains/Task"
	"go-echo-rest-api-sample/src/Presentations/Requests/Task"
	"go-echo-rest-api-sample/src/Shared"
	"time"
)

type SaveTaskUseCase struct {
	taskRepository DomainServices.ITaskRepository
}

func NewSaveTaskUseCase(taskRepository DomainServices.ITaskRepository) *SaveTaskUseCase {
	return &SaveTaskUseCase{
		taskRepository: taskRepository,
	}
}

func (u *SaveTaskUseCase) Execute(request Requests.SaveTaskRequest) error {
	t, err := time.Parse(Domains.DeadlineFormat, request.Deadline)
	if err != nil {
		return Shared.NewSampleError("001-001", "締切日のフォーマットが不正です")
	}

	task, err := Domains.CreateNewTask(request.Name, t)
	if err != nil {
		return err
	}
	if err := u.taskRepository.SaveTask(task); err != nil {
		return err
	}

	return nil
}
