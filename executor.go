package libgin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/rbac"
)

type ExecutorOnOpenFunc func(c *gin.Context) error
type ExecutorOnCloseFunc func()
type ExecutorOnErrorFunc func(err error)
type ExecutorOnPanicFunc func(x any)
type ExecutorOnSendFunc func(err error)

type ExecutorOnTaskFunc func() error

////////////////////////////////////////////////////////////////////////////////

// Executor 是默认的 REST 请求执行对象；
// 它为 REST API 提供一个便捷的可编程执行过程。
type Executor struct {

	// attrs

	Responder Responder
	Context   *gin.Context

	Body1  any
	Body2  any
	Error  error
	Status int

	// callbacks
	OnOpen  ExecutorOnOpenFunc
	OnSend  ExecutorOnSendFunc
	OnTask  ExecutorOnTaskFunc
	OnError ExecutorOnErrorFunc
	OnPanic ExecutorOnPanicFunc
	OnClose ExecutorOnCloseFunc
}

// func (inst *Executor) Init(f1 ExecutorOnOpenFunc, f2 ExecutorOnTaskFunc) {
// 	inst.OnOpen = f1
// 	inst.OnTask = f2
// }

func (inst *Executor) Execute() {

	fnOpen := inst.OnOpen
	fnSend := inst.OnSend
	fnTask := inst.OnTask
	fnError := inst.OnError
	fnPanic := inst.OnPanic
	fnClose := inst.OnClose

	////////////////////////

	fnOpen = inst.innerPrepareOnOpen(fnOpen)
	fnTask = inst.innerPrepareOnTask(fnTask)
	fnSend = inst.innerPrepareOnSend(fnSend)
	fnError = inst.innerPrepareOnError(fnError)
	fnPanic = inst.innerPrepareOnPanic(fnPanic)
	fnClose = inst.innerPrepareOnClose(fnClose)

	////////////////////////

	inst.OnSend = fnSend

	////////////////////////

	ctx := inst.Context
	err := fnOpen(ctx)
	defer fnClose()

	defer func() {
		x := recover()
		fnPanic(x)
	}()

	if err == nil {
		err = fnTask()
	}

	fnError(err)
}

//// prepare functions /////////////////////

func (inst *Executor) innerPrepareOnOpen(fn ExecutorOnOpenFunc) ExecutorOnOpenFunc {
	if fn == nil {
		fn = inst.DefaultOnOpen
	}
	return fn
}

func (inst *Executor) innerPrepareOnClose(fn ExecutorOnCloseFunc) ExecutorOnCloseFunc {
	if fn == nil {
		fn = inst.DefaultOnClose
	}
	return fn
}

func (inst *Executor) innerPrepareOnSend(fn ExecutorOnSendFunc) ExecutorOnSendFunc {
	if fn == nil {
		fn = inst.DefaultOnSend
	}
	return fn
}

func (inst *Executor) innerPrepareOnTask(fn ExecutorOnTaskFunc) ExecutorOnTaskFunc {
	if fn == nil {
		fn = inst.DefaultOnTask
	}
	return fn
}

func (inst *Executor) innerPrepareOnError(fn ExecutorOnErrorFunc) ExecutorOnErrorFunc {
	if fn == nil {
		fn = inst.DefaultOnError
	}
	return fn
}

func (inst *Executor) innerPrepareOnPanic(fn ExecutorOnPanicFunc) ExecutorOnPanicFunc {
	if fn == nil {
		fn = inst.DefaultOnPanic
	}
	return fn
}

////  default functions ///////////////////////

func (inst *Executor) DefaultOnOpen(c *gin.Context) error {
	return nil // a NOP open
}

func (inst *Executor) DefaultOnTask() error {
	return nil // a NOP task
}

func (inst *Executor) DefaultOnClose() {
	err := inst.Error
	sender := inst.OnSend
	sender(err)
}

func (inst *Executor) DefaultOnError(err error) {
	if err == nil {
		return
	}
	inst.Error = err
}

func (inst *Executor) DefaultOnPanic(x any) {

	if x == nil {
		return
	}

	err, ok := x.(error)
	if ok {
		inst.DefaultOnError(err)
		return
	}

	str, ok := x.(string)
	if ok {
		err = fmt.Errorf("panic: %s", str)
		inst.DefaultOnError(err)
		return
	}

	err = fmt.Errorf("panic: %v", x)
	inst.DefaultOnError(err)
}

func (inst *Executor) DefaultOnSend(err error) {

	sender := inst.Responder
	ctx := inst.Context
	b2 := inst.Body2
	code := inst.Status

	if err != nil {
		code = http.StatusInternalServerError
	}
	code = inst.innerComputeStatusCode(code, b2)

	resp := new(Response)
	resp.Context = ctx
	resp.Data = b2
	resp.Error = err
	resp.Responder = ""
	resp.Status = code

	sender.Send(resp)
}

func (inst *Executor) innerComputeStatusCode(code int, body2 any) int {
	ref, ok := body2.(rbac.VORef)
	if ok {
		b2vo := ref.GetTarget()
		c1 := code
		c2 := b2vo.Status
		if c1 < c2 {
			code = c2
		} else {
			code = c1
		}
		if code == 0 {
			code = http.StatusOK
		}
		b2vo.Status = code
	}
	return code
}
