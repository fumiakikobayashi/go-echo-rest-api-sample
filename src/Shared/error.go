package Shared

import (
	"fmt"
	"go.uber.org/zap"
)

type MyError struct {
	Code       string
	Message    string
	StackTrace string
}

func (me *MyError) Error() string {
	return fmt.Sprintf("my error: code[%s], message[%s]", me.Code, me.Message)
}

func New(code string, message string) *MyError {
	stack := zap.Stack("").String
	return &MyError{
		Code:       code,
		Message:    message,
		StackTrace: stack,
	}
}
