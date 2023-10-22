package libgin

import "github.com/gin-gonic/gin"

// Response 表示一个 REST 响应结果
type Response struct {
	Context   *gin.Context
	Status    int
	Error     error
	Responder string // 用于选择 Responder，它的值为 ResponderRegistration.Name
	Data      any
}

// Responder 是 REST 响应发送器的接口
type Responder interface {
	// ResponderRegistry

	Accept(resp *Response) bool

	Send(resp *Response)
}

// ResponderRegistration 是 Responder 的注册信息
type ResponderRegistration struct {
	Priority  int // 优先级，数值越大，越先处理
	Name      string
	Enabled   bool
	Responder Responder
}

// ResponderRegistry 是 Responder 的注册接口
type ResponderRegistry interface {
	ListRegistrations() []*ResponderRegistration
}
