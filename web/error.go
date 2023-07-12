package web

import (
	"fmt"
	"net/http"
)

// Error 表示一个携带了 HTTP 状态的错误
type Error interface {
	error

	Status() int
}

// NewError 新建一个携带了 HTTP 状态的错误
func NewError(status int, format string, args ...any) Error {
	msg := fmt.Sprintf(format, args...)
	return &innerWebError{status: status, msg: msg}
}

////////////////////////////////////////////////////////////////////////////////

type innerWebError struct {
	msg    string
	status int
}

func (inst *innerWebError) Error() string {
	return inst.msg
}

func (inst *innerWebError) String() string {
	code := inst.status
	text := http.StatusText(code)
	return fmt.Sprintf("HTTP %d %s (%s)", code, text, inst.msg)
}

func (inst *innerWebError) Status() int {
	return inst.status
}
