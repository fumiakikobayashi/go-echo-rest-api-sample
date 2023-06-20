package UseCases

import (
	"fmt"
	Requests "go-ddd-rest-api-sample/src/Presentations/Requests/Task"
	"go-ddd-rest-api-sample/src/Shared"
	Dto "go-ddd-rest-api-sample/src/UseCases/Dto/Task"
	uShared "go-ddd-rest-api-sample/src/UseCases/Shared"
)

type GetTasksUseCase struct {
	taskRepository TaskRepositoryInterface
	logger         Shared.LoggerInterface
}

func NewGetTasksUseCase(taskRepository TaskRepositoryInterface, logger Shared.LoggerInterface) *GetTasksUseCase {
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
		return Dto.TaskListDto{}, fmt.Errorf("タスク一覧の取得に失敗しました")
	}

	taskListDto, err := CreateTaskDtoList(taskList)
	if err != nil {
		return Dto.TaskListDto{}, fmt.Errorf("DTOの作成に失敗しました")
	}
	return taskListDto, nil
}
