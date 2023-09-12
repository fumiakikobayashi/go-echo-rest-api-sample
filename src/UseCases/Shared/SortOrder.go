package Shared

import (
	"go-echo-rest-api-sample/src/Shared"
)

type SortOrder string

const (
	Asc  SortOrder = "asc"
	Desc SortOrder = "desc"
)

func NewSortOrder(sortOrder string) (SortOrder, error) {
	switch sortOrder {
	case string(Asc):
		return Asc, nil
	case string(Desc):
		return Desc, nil
	case "":
		return Desc, nil
	default:
		return "", Shared.NewSampleError("001-001", "想定しないSortOrderが入力されました")
	}
}

func (so SortOrder) GetValue() string {
	return string(so)
}
