package UseCases

import (
	"fmt"
	Domains "go-ddd-rest-api-sample/src/Domains/Task"
	"go-ddd-rest-api-sample/src/Presentations/Requests"
	Dto "go-ddd-rest-api-sample/src/UseCases/Dto/Task"
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
	taskId, _ := Domains.NewTaskId(request.TaskId)
	task, err := u.taskRepository.GetTask(taskId)
	if err != nil {
		return Dto.TaskDto{}, fmt.Errorf("タスクの取得に失敗しました")
	}

	taskDto, err := CreateTaskDto(task)
	if err != nil {
		return Dto.TaskDto{}, fmt.Errorf("タスクDTOの生成に失敗しました")
	}
	return taskDto, nil
}
