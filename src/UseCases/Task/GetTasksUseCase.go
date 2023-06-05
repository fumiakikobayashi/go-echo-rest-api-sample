package UseCases

import (
	"fmt"
	Dto "go-ddd-rest-api-sample/src/UseCases/Dto/Task"
)

type GetTasksUseCase struct {
	taskRepository TaskRepositoryInterface
}

func NewGetTasksUseCase(taskRepository TaskRepositoryInterface) *GetTasksUseCase {
	return &GetTasksUseCase{
		taskRepository: taskRepository,
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
