package UseCases

import (
	"go-echo-rest-api-sample/src/DomainServices"
	"go-echo-rest-api-sample/src/Presentations/Requests/Task"
	"go-echo-rest-api-sample/src/UseCases/Dto/Task"
	uShared "go-echo-rest-api-sample/src/UseCases/Shared"
)

type GetTasksUseCase struct {
	taskRepository DomainServices.ITaskRepository
}

func NewGetTasksUseCase(taskRepository DomainServices.ITaskRepository) *GetTasksUseCase {
	return &GetTasksUseCase{
		taskRepository: taskRepository,
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
