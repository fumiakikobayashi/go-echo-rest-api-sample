package UseCases

import (
	Domains "go-ddd-rest-api-sample/src/Domains/Task"
	Requests "go-ddd-rest-api-sample/src/Presentations/Requests/Task"
	"go-ddd-rest-api-sample/src/Shared"
	"go-ddd-rest-api-sample/src/Shared/Errors"
)

type DeleteTaskUseCase struct {
	taskRepository TaskRepositoryInterface
	logger         Shared.LoggerInterface
}

func NewDeleteTaskUseCase(taskRepository TaskRepositoryInterface, logger Shared.LoggerInterface) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{
		taskRepository: taskRepository,
		logger:         logger,
	}
}

func (u *DeleteTaskUseCase) Execute(request Requests.DeleteTaskRequest) error {
	taskId, err := Domains.NewTaskId(request.TaskId)
	if err != nil {
		return err
	}

	task, err := u.taskRepository.GetTask(taskId)
	if err != nil {
		return err
	}
	if task == nil {
		return Errors.New("001-001", "指定されたタスクが存在しません")
	}

	if err := u.taskRepository.DeleteTask(taskId); err != nil {
		return err
	}
	return nil
}
