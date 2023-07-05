package UseCases

import (
	Domains "go-ddd-rest-api-sample/src/Domains/Task"
	Requests "go-ddd-rest-api-sample/src/Presentations/Requests/Task"
	"go-ddd-rest-api-sample/src/Shared"
	Dto "go-ddd-rest-api-sample/src/UseCases/Dto/Task"
)

type GetTaskUseCase struct {
	taskRepository TaskRepositoryInterface
	logger         *Shared.LoggerInterface
}

func NewGetTaskUseCase(taskRepository TaskRepositoryInterface, logger *Shared.LoggerInterface) *GetTaskUseCase {
	return &GetTaskUseCase{
		taskRepository: taskRepository,
		logger:         logger,
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
