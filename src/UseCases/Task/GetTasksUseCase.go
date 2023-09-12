package UseCases

import (
	Requests "go-echo-rest-api-sample/src/Presentations/Requests/Task"
	"go-echo-rest-api-sample/src/Shared"
	Dto "go-echo-rest-api-sample/src/UseCases/Dto/Task"
	uShared "go-echo-rest-api-sample/src/UseCases/Shared"
)

type GetTasksUseCase struct {
	taskRepository TaskRepositoryInterface
	logger         *Shared.LoggerInterface
}

func NewGetTasksUseCase(taskRepository TaskRepositoryInterface, logger *Shared.LoggerInterface) *GetTasksUseCase {
	return &GetTasksUseCase{
		taskRepository: taskRepository,
		logger:         logger,
	}
}

func (u *GetTasksUseCase) Execute(tasksRequest Requests.GetTasksRequest) (Dto.TaskListDto, error) {
	sortType, err := uShared.NewSortType(tasksRequest.Sort)
	if err != nil {
		return Dto.TaskListDto{}, err
	}
	sortOrder, err := uShared.NewSortOrder(tasksRequest.Order)
	if err != nil {
		return Dto.TaskListDto{}, err
	}

	taskList, err := u.taskRepository.GetTasks(sortType, sortOrder)
	if err != nil {
		return Dto.TaskListDto{}, err
	}

	return CreateTaskDtoList(taskList), nil
}
