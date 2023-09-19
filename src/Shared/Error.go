package Shared

type SampleError struct {
	Code    string
	Message string
}

func (me *SampleError) Error() string {
	return me.Message
}

func NewSampleError(code string, message string) *SampleError {
	return &SampleError{
		Code:    code,
		Message: message,
	}
}
