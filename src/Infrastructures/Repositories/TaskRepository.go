package Repositories

import (
	"github.com/jinzhu/gorm"
	Domains "go-ddd-rest-api-sample/src/Domains/Task"
	"go-ddd-rest-api-sample/src/Infrastructures/Models"
	UseCase "go-ddd-rest-api-sample/src/UseCases/Task"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) UseCase.TaskRepositoryInterface {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) GetTasks() (*Domains.TaskList, error) {
	var taskModels []Models.TaskModel
	taskList := Domains.NewTaskList()

	if err := r.db.Table("tasks").Find(&taskModels).Error; err != nil {
		return taskList, err
	}

	for _, taskModel := range taskModels {
		task, _ := Domains.CreateTask(taskModel)
		if err := taskList.Push(task); err != nil {
			return taskList, err
		}
	}

	return taskList, nil
}

func (r *taskRepository) GetTask(taskId Domains.TaskId) (*Domains.Task, error) {
	var taskModel Models.TaskModel

	if err := r.db.Table("tasks").First(&taskModel, taskId.GetValue()).Error; err != nil {
		return &Domains.Task{}, err
	}

	task, err := Domains.CreateTask(taskModel)
	if err != nil {
		return &Domains.Task{}, err
	}
	return task, nil
}
