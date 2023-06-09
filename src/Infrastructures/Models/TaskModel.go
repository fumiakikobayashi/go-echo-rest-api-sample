package Models

import "time"

type TaskModel struct {
	ID          int `gorm:"primary_key"`
	Name        string
	Deadline    time.Time
	IsFavorite  bool
	IsCompleted bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t TaskModel) TableName() string {
	return "tasks"
}
