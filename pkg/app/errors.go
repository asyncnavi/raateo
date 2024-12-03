package app

import (
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	msg     string
	code    string
	status  int
	details map[string]interface{}
}

var _ error = new(AppError)

func (a *AppError) Code() string {
	return a.code
}

func (a *AppError) Error() string {
	return a.msg
}

func (a *AppError) Status() int {
	return a.status
}

func (a *AppError) Details() map[string]interface{} {
	return a.details
}

func (a *AppError) WithCode(code string) *AppError {
	a.code = code
	return a
}

func (a *AppError) WithStatus(status int) *AppError {
	a.status = status
	return a
}

func (a *AppError) WithAttribute(key string, value interface{}) *AppError {
	a.details[key] = value
	return a
}

func Error(msg string) *AppError {
	return &AppError{msg: msg, status: http.StatusBadRequest}
}

func NewError(msg string, code string, status int) *AppError {
	return &AppError{msg: msg, code: code, status: status}
}

func Errorf(msg string, vars ...any) *AppError {
	return &AppError{msg: fmt.Sprintf(msg, vars...), status: http.StatusBadRequest}
}

func (a *AppError) Rule(rule string, fields ...string) *AppError {
	code := rule
	if len(fields) != 0 {
		code = fmt.Sprintf("%s.%s", rule, strings.Join(fields, "."))
	}
	if len(a.code) == 0 {
		a.code = code
	} else {
		a.code = fmt.Sprintf("%s.%s", a.code, code)
	}
	return a
}

func (a *AppError) Domain() *AppError {
	a.code = "dmn"
	return a
}

func (a *AppError) NotUnique(fields ...string) *AppError {
	return a.Rule("not_unique", fields...)
}

func (a *AppError) NotAvailable(fields ...string) *AppError {
	return a.Rule("not_available", fields...)
}

func (a *AppError) Validation(fields ...string) *AppError {
	return a.Rule("vldn", fields...)
}
