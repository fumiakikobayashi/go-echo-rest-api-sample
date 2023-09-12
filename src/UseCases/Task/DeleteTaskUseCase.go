package UseCases

import (
	Domains "go-echo-rest-api-sample/src/Domains/Task"
	Requests "go-echo-rest-api-sample/src/Presentations/Requests/Task"
	"go-echo-rest-api-sample/src/Shared"
)

type DeleteTaskUseCase struct {
	taskRepository TaskRepositoryInterface
	logger         *Shared.LoggerInterface
}

func NewDeleteTaskUseCase(taskRepository TaskRepositoryInterface, logger *Shared.LoggerInterface) *DeleteTaskUseCase {
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
		return Shared.NewSampleError("001-001", "指定されたタスクが存在しません")
	}

	if err := u.taskRepository.DeleteTask(taskId); err != nil {
		return err
	}
	return nil
}
