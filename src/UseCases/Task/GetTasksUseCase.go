package UseCases

import (
	"fmt"
	"go-ddd-rest-api-sample/sdk"
	Dto "go-ddd-rest-api-sample/src/UseCases/Dto/Task"
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

func (u *GetTasksUseCase) Execute() (Dto.TaskListDto, error) {
	taskList, err := u.taskRepository.GetTasks()
	if err != nil {
		return Dto.TaskListDto{}, fmt.Errorf("タスク一覧の取得に失敗しました")
	}

	taskListDto, err := CreateTaskDtoList(taskList)
	if err != nil {
		return Dto.TaskListDto{}, fmt.Errorf("タスク一覧の取得に失敗しました")
	}
	return taskListDto, nil
}
