package errorx

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
)

type Error struct {
	Code    int32
	Message string
}

func (e *Error) Error() string {
	s, _ := json.Marshal(e)
	return string(s)
}

func New(msg string) *Error {
	return &Error{
		Code:    1000,
		Message: msg,
	}
}

func NewWithCode(code int32, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}

func NewF(format string, args ...interface{}) *Error {
	return &Error{
		Code:    1000,
		Message: fmt.Sprintf(format, args...),
	}
}

func NewFWithCode(code int32, format string, args ...interface{}) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}
