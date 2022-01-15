package errors

import (
	"fmt"
	"net/http"
)

type Error struct {
	// Code 业务代码
	Code int64 `json:"code"`
	// Status http状态码
	Status int `json:"-"`
	// Message 错误提示消息
	Message interface{} `json:"message"`
}

func New(code int64, status int, message interface{}) *Error {
	return &Error{
		Code:    code,
		Status:  status,
		Message: message,
	}
}

func (e *Error) Error() string {
	switch m := e.Message.(type) {
	case string:
		return m
	case map[string]string:
		var msg string
		for k, v := range m {
			msg += fmt.Sprintf("%s : %s", k, v)
		}
		return msg
	default:
		return ""
	}
}

func (e *Error) GetMessage() string {
	if e != nil {
		return e.Error()
	}
	return ""
}

func (e *Error) GetCode() int64 {
	if e != nil {
		return e.Code
	}
	return 0
}

func (e *Error) GetStatus() int {
	if e != nil {
		return e.Status
	}
	return http.StatusOK
}
