package Domains

import "errors"

type UserId struct {
	Value int
}

const MIN_VALUE = 0

func NewUserId(value int) (UserId, error) {
	if value < MIN_VALUE {
		errors.New("UserIdは1以上の数値を入力する必要があります")
	}
	return UserId{Value: value}, nil
}

func (userId *UserId) GetValue() int {
	return userId.Value
}
