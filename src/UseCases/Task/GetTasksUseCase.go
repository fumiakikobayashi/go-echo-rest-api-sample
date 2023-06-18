package UseCases

import (
	"fmt"
	"go-ddd-rest-api-sample/sdk"
	Requests "go-ddd-rest-api-sample/src/Presentations/Requests/Task"
	Dto "go-ddd-rest-api-sample/src/UseCases/Dto/Task"
	"go-ddd-rest-api-sample/src/UseCases/Shared"
)

type GetTasksUseCase struct {
	taskRepository TaskRepositoryInterface
	logger         sdk.LoggerInterface
}

func NewGetTasksUseCase(taskRepository TaskRepositoryInterface, logger sdk.LoggerInterface) *GetTasksUseCase {
	return &GetTasksUseCase{
		taskRepository: taskRepository,
		logger:         logger,
	}
}

func (u *GetTasksUseCase) Execute(tasksRequest Requests.GetTasksRequest) (Dto.TaskListDto, error) {
	sortType, err := Shared.NewSortType(tasksRequest.Sort)
	if err != nil {
		return Dto.TaskListDto{}, err
	}
	sortOrder, err := Shared.NewSortOrder(tasksRequest.Order)
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
