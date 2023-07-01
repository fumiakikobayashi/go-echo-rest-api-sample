package Shared

import (
	"go.uber.org/zap"
)

type SampleError struct {
	Code       string
	Message    string
	StackTrace string
}

func (me *SampleError) Error() string {
	return me.Message
}

func NewSampleError(code string, message string) *SampleError {
	stack := zap.Stack("").String
	return &SampleError{
		Code:       code,
		Message:    message,
		StackTrace: stack,
	}
}
