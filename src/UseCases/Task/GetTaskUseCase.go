package UseCases

import (
	"go-echo-rest-api-sample/src/Domains/Task"
	"go-echo-rest-api-sample/src/Presentations/Requests/Task"
	"go-echo-rest-api-sample/src/UseCases/Dto/Task"
)

type GetTaskUseCase struct {
	taskRepository TaskRepositoryInterface
}

func NewGetTaskUseCase(taskRepository TaskRepositoryInterface) *GetTaskUseCase {
	return &GetTaskUseCase{
		taskRepository: taskRepository,
	}
}

func (u *GetTaskUseCase) Execute(request Requests.GetTaskRequest) (Dto.TaskDto, error) {
	taskId, err := Domains.NewTaskId(request.TaskId)
	if err != nil {
		return Dto.TaskDto{}, err
	}
	task, err := u.taskRepository.GetTask(taskId)
	if err != nil {
		return Dto.TaskDto{}, err
	}

	return CreateTaskDto(task), nil
}
