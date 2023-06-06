package UseCases

import (
	"fmt"
	Domains "go-ddd-rest-api-sample/src/Domains/Task"
	"go-ddd-rest-api-sample/src/Presentations/Requests"
)

type DeleteTaskUseCase struct {
	taskRepository TaskRepositoryInterface
}

func NewDeleteTaskUseCase(taskRepository TaskRepositoryInterface) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{
		taskRepository: taskRepository,
	}
}

func (u *DeleteTaskUseCase) Execute(request Requests.DeleteTaskRequest) error {
	taskId, err := Domains.NewTaskId(request.TaskId)
	if err != nil {
		return err
	}

	if err := u.taskRepository.DeleteTask(taskId); err != nil {
		return fmt.Errorf("タスクの削除に失敗しました")
	}
	return nil
}
