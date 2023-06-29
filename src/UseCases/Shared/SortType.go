package Shared

import (
	"go-ddd-rest-api-sample/src/Shared"
)

type SortType string

const (
	Id       SortType = "id"
	Name     SortType = "name"
	Deadline SortType = "deadline"
	Favorite SortType = "favorite"
)

func NewSortType(sortType string) (SortType, error) {
	switch sortType {
	case string(Name):
		return Name, nil
	case string(Deadline):
		return Deadline, nil
	case string(Favorite):
		return Favorite, nil
	case "":
		return Id, nil
	default:
		return "", Shared.New("001-001", "想定しないSortTypeが入力されました")
	}
}

func (so SortType) GetValue() string {
	return string(so)
}
