package UseCases

import (
	"fmt"
	"go-ddd-rest-api-sample/sdk"
	Domains "go-ddd-rest-api-sample/src/Domains/Task"
	Requests "go-ddd-rest-api-sample/src/Presentations/Requests/Task"
	Dto "go-ddd-rest-api-sample/src/UseCases/Dto/Task"
)

type GetTaskUseCase struct {
	taskRepository TaskRepositoryInterface
	logger         sdk.LoggerInterface
}

func NewGetTaskUseCase(taskRepository TaskRepositoryInterface, logger sdk.LoggerInterface) *GetTaskUseCase {
	return &GetTaskUseCase{
		taskRepository: taskRepository,
		logger:         logger,
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
