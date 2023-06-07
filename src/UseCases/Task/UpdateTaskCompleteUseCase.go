package UseCases

import (
	"go-ddd-rest-api-sample/sdk"
	Domains "go-ddd-rest-api-sample/src/Domains/Task"
	Requests "go-ddd-rest-api-sample/src/Presentations/Requests/Task"
)

type UpdateTaskCompleteUseCase struct {
	taskRepository TaskRepositoryInterface
	logger         sdk.LoggerInterface
}

func NewUpdateTaskCompleteUseCase(taskRepository TaskRepositoryInterface, logger sdk.LoggerInterface) *UpdateTaskCompleteUseCase {
	return &UpdateTaskCompleteUseCase{
		taskRepository: taskRepository,
		logger:         logger,
	}
}

func (u *UpdateTaskCompleteUseCase) Execute(request Requests.UpdateTaskCompleteRequest) error {
	taskId, err := Domains.NewTaskId(request.TaskId)
	if err != nil {
		return err
	}

	task, err := u.taskRepository.GetTask(taskId)
	if err != nil {
		return err
	}

	task.UpdateTaskComplete()
	if err := u.taskRepository.UpdateTask(task); err != nil {
		return err
	}
	return nil
}
