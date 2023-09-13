package UseCases

import (
	"go-echo-rest-api-sample/src/Domains/Task"
	"go-echo-rest-api-sample/src/Presentations/Requests/Task"
)

type UpdateTaskCompleteUseCase struct {
	taskRepository TaskRepositoryInterface
}

func NewUpdateTaskCompleteUseCase(taskRepository TaskRepositoryInterface) *UpdateTaskCompleteUseCase {
	return &UpdateTaskCompleteUseCase{
		taskRepository: taskRepository,
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
