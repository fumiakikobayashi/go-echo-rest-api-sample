package UseCases

import (
	"go-echo-rest-api-sample/src/Presentations/Requests/Task"
	"go-echo-rest-api-sample/src/UseCases/Dto/Task"
	Shared2 "go-echo-rest-api-sample/src/UseCases/Shared"
)

type GetTasksUseCase struct {
	taskRepository TaskRepositoryInterface
}

func NewGetTasksUseCase(taskRepository TaskRepositoryInterface) *GetTasksUseCase {
	return &GetTasksUseCase{
		taskRepository: taskRepository,
	}
}

func (u *GetTasksUseCase) Execute(tasksRequest Requests.GetTasksRequest) (Dto.TaskListDto, error) {
	sortType, err := Shared2.NewSortType(tasksRequest.Sort)
	if err != nil {
		return Dto.TaskListDto{}, err
	}
	sortOrder, err := Shared2.NewSortOrder(tasksRequest.Order)
	if err != nil {
		return Dto.TaskListDto{}, err
	}

	taskList, err := u.taskRepository.GetTasks(sortType, sortOrder)
	if err != nil {
		return Dto.TaskListDto{}, err
	}

	return CreateTaskDtoList(taskList), nil
}
