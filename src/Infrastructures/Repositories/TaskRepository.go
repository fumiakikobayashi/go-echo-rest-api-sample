package Repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-ddd-rest-api-sample/sdk"
	Domains "go-ddd-rest-api-sample/src/Domains/Task"
	"go-ddd-rest-api-sample/src/Infrastructures/Models"
	"go-ddd-rest-api-sample/src/UseCases/Shared"
	UseCase "go-ddd-rest-api-sample/src/UseCases/Task"
	"time"
)

type taskRepository struct {
	db     *gorm.DB
	logger sdk.LoggerInterface
}

func NewTaskRepository(db *gorm.DB, logger sdk.LoggerInterface) UseCase.TaskRepositoryInterface {
	return &taskRepository{
		db:     db,
		logger: logger,
	}
}

func (r *taskRepository) GetTasks(sortType Shared.SortType, sortOrder Shared.SortOrder) (*Domains.TaskList, error) {
	var taskModels []Models.TaskModel
	var sortColumn string
	taskList := Domains.NewTaskList()

	switch sortType {
	case Shared.Name:
		sortColumn = "name"
	case Shared.Deadline:
		sortColumn = "deadline"
	case Shared.Favorite:
		sortColumn = "is_favorite"
	default:
		sortColumn = "id"
	}

	if err := r.db.Table("tasks").Order(fmt.Sprintf("%s %s", sortColumn, sortOrder.GetValue())).Find(&taskModels).Error; err != nil {
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

func (r *taskRepository) SaveTask(task *Domains.Task) error {
	isFavorite := task.GetIsFavorite()
	isCompleted := task.GetIsCompleted()

	if err := r.db.Table("tasks").Create(&Models.TaskModel{
		ID:          task.GetTaskId().GetValue(),
		Name:        task.GetName(),
		Deadline:    task.GetDeadline(),
		IsFavorite:  &isFavorite,
		IsCompleted: &isCompleted,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) UpdateTask(task *Domains.Task) error {
	taskId := task.GetTaskId().GetValue()
	isFavorite := task.GetIsFavorite()
	isCompleted := task.GetIsCompleted()

	if err := r.db.Table("tasks").Where("id = ?", taskId).Updates(&Models.TaskModel{
		ID:          taskId,
		Name:        task.GetName(),
		Deadline:    task.GetDeadline(),
		IsFavorite:  &isFavorite,
		IsCompleted: &isCompleted,
		UpdatedAt:   time.Now(),
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) DeleteTask(taskId Domains.TaskId) error {
	var taskModel Models.TaskModel

	if err := r.db.Table("tasks").Where("id = ?", taskId.GetValue()).Delete(&taskModel).Error; err != nil {
		return err
	}

	return nil
}
