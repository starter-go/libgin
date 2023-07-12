package libgin

import "github.com/gin-gonic/gin"

// Response 表示一个 REST 响应结果
type Response struct {
	Context *gin.Context
	Status  int
	Error   error
	Data    any
}

// Responder 是 REST 响应发送器的接口
type Responder interface {
	Send(resp *Response)
}
