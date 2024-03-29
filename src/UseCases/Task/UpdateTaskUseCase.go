package UseCases

import (
	"go-echo-rest-api-sample/src/DomainServices"
	"go-echo-rest-api-sample/src/Domains/Task"
	"go-echo-rest-api-sample/src/Presentations/Requests/Task"
	"go-echo-rest-api-sample/src/Shared"
	"time"
)

type UpdateTaskUseCase struct {
	taskRepository DomainServices.ITaskRepository
}

func NewUpdateTaskUseCase(taskRepository DomainServices.ITaskRepository) *UpdateTaskUseCase {
	return &UpdateTaskUseCase{
		taskRepository: taskRepository,
	}
}

func (u *UpdateTaskUseCase) Execute(request Requests.UpdateTaskRequest) error {
	taskId, err := Domains.NewTaskId(request.TaskId)
	if err != nil {
		return err
	}
	task, err := u.taskRepository.GetTask(taskId)
	if err != nil {
		return err
	}

	t, err := time.Parse(Domains.DeadlineFormat, request.Deadline)
	if err != nil {
		return Shared.NewSampleError("001-001", "締切日のフォーマットが不正です")
	}
	task.UpdateTask(request.Name, t)

	if err := u.taskRepository.UpdateTask(task); err != nil {
		return err
	}
	return nil
}
