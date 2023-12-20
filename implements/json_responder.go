package implements

import (
	"net/http"
	"time"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgin"
	"github.com/starter-go/libgin/web"
	"github.com/starter-go/rbac"
)

// JSONGinResponder ...
type JSONGinResponder struct {

	//starter:component
	_as func(libgin.ResponderRegistry, libgin.Responder) //starter:as(".",".")
}

func (inst *JSONGinResponder) _impl() (libgin.ResponderRegistry, libgin.Responder) {
	return inst, inst
}

// ListRegistrations ...
func (inst *JSONGinResponder) ListRegistrations() []*libgin.ResponderRegistration {
	r1 := &libgin.ResponderRegistration{
		Enabled:   true,
		Priority:  0,
		Name:      "security-gin-rbac-responder",
		Responder: inst,
	}
	return []*libgin.ResponderRegistration{r1}
}

// Accept ...
func (inst *JSONGinResponder) Accept(resp *libgin.Response) bool {
	return true
}

// Send ...
func (inst *JSONGinResponder) Send(resp *libgin.Response) {

	now := time.Now()
	ctx := resp.Context
	err := resp.Error
	data := resp.Data
	status := resp.Status

	werr, ok := err.(web.Error)
	if ok {
		status = werr.Status()
	}

	if status == 0 {
		if err == nil {
			status = http.StatusOK
		} else {
			status = http.StatusInternalServerError
		}
	}

	vg, ok := data.(rbac.VOGetter)
	if ok {
		v := vg.GetVO()
		if err != nil {
			v.Error = err.Error()
		}
		v.Message = http.StatusText(status)
		v.Status = status
		v.Time = now
		v.Timestamp = lang.NewTime(now)
	}

	ctx.JSON(status, data)
}

// import (
// 	"net/http"
// 	"time"

// 	"github.com/starter-go/base/lang"
// 	"github.com/starter-go/libgin"
// 	"github.com/starter-go/libgin/web"
// )

// // RESTfulResponder 这个实现仅供参考，实际应该使用 “github.com/starter-go/module-gin-security” 模块里实现的组件！
// // RESTfulResponder 默认的 RESTful 响应发送器
// type RESTfulResponder struct {
// 	// ## starter:component
// 	_as func(libgin.Responder) //starter:as(".")
// }

// func (inst *RESTfulResponder) _impl() {
// 	inst._as(inst)
// }

// // Registration 返回注册信息
// func (inst *RESTfulResponder) Registration() *libgin.ResponderRegistration {
// 	return nil
// }

// // Accept 确认是否接受传入的响应内容
// func (inst *RESTfulResponder) Accept(resp *libgin.Response) bool {
// 	return false
// }

// // Send 发送响应
// func (inst *RESTfulResponder) Send(resp *libgin.Response) {

// 	c := resp.Context
// 	data := resp.Data
// 	err := resp.Error
// 	status := resp.Status
// 	msg := ""
// 	errStr := ""

// 	if err != nil {
// 		we, ok := err.(web.Error)
// 		if ok && status == 0 {
// 			status = we.Status()
// 		}
// 		errStr = err.Error()
// 	}

// 	if status == 0 {
// 		if err == nil {
// 			status = http.StatusOK
// 		} else {
// 			status = http.StatusInternalServerError
// 		}
// 	}
// 	msg = http.StatusText(status)

// 	baseGetter, ok := data.(JSONRootGetter)
// 	if ok {
// 		vo := baseGetter.GetRoot()
// 		if vo.Message == "" {
// 			vo.Message = msg
// 		}
// 		if vo.Error == "" {
// 			vo.Error = errStr
// 		}
// 		now := time.Now()
// 		vo.Status = status
// 		vo.Time = now
// 		vo.Timestamp = lang.NewTime(now)
// 	}

// 	c.JSON(status, data)
// }

// ////////////////////////////////////////////////////////////////////////////////

// // JSONRoot 表示基本的 JSON VO 对象
// type JSONRoot struct {
// 	Message   string
// 	Error     string
// 	Status    int
// 	Time      time.Time
// 	Timestamp lang.Time
// }

// // JSONRootGetter 是获取 JSONRoot 的接口
// type JSONRootGetter interface {
// 	GetRoot() *JSONRoot
// }
