package utils

import (
	"fmt"
)

type ResponseError struct {
	Message string
	Code    int
}

func NewResponseError(msg string, code int) *ResponseError {
	return &ResponseError{
		Message: msg,
		Code:    code,
	}
}

func (r ResponseError) Error() string {
	return fmt.Sprintf("Code: %v | Error: %v", r.Code, r.Message)
}

