package implements

import (
	"net/http"
	"time"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgin"
	"github.com/starter-go/libgin/web"
	"github.com/starter-go/rest"
)

// RESTfulResponder 默认的 RESTful 响应发送器
type RESTfulResponder struct {
	//starter:component
	_as func(libgin.Responder) //starter:as("#")
}

func (inst *RESTfulResponder) _impl() {
	inst._as(inst)
}

// Send 发送响应
func (inst *RESTfulResponder) Send(resp *libgin.Response) {

	c := resp.Context
	data := resp.Data
	err := resp.Error
	status := resp.Status
	msg := ""
	errStr := ""

	if err != nil {
		we, ok := err.(web.Error)
		if ok && status == 0 {
			status = we.Status()
		}
		errStr = err.Error()
	}

	if status == 0 {
		if err == nil {
			status = http.StatusOK
		} else {
			status = http.StatusInternalServerError
		}
	}
	msg = http.StatusText(status)

	baseGetter, ok := data.(rest.BaseGetter)
	if ok {
		vo := baseGetter.GetBase()
		if vo.Message == "" {
			vo.Message = msg
		}
		if vo.Error == "" {
			vo.Error = errStr
		}
		now := time.Now()
		vo.Status = status
		vo.Time = now
		vo.Timestamp = lang.NewTime(now)
	}

	c.JSON(status, data)
}
