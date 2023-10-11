package libgin

import "github.com/gin-gonic/gin"

// Controller 表示一个 gin 控制器
type Controller interface {
	Registration() *ControllerRegistration
}

// ControllerRegistration 是控制器注册信息
type ControllerRegistration struct {
	Groups []string                   // 分组名称, 以 ',' 号分隔多个名称，如果为空表示默认分组
	Route  func(rp RouterProxy) error // 控制器
}

// RouterProxy ...
type RouterProxy interface {
	For(path string) RouterProxy
	Route(r *Routing)
	Handle(method string, path string, h ...gin.HandlerFunc)
	NoMethod(h ...gin.HandlerFunc)
	NoRoute(h ...gin.HandlerFunc)

	GET(path string, h ...gin.HandlerFunc)
	POST(path string, h ...gin.HandlerFunc)
	PUT(path string, h ...gin.HandlerFunc)
	DELETE(path string, h ...gin.HandlerFunc)
}

// Routing 表示一条路由记录
type Routing struct {
	Method     string
	Path       string
	Priority   int
	NoMethod   bool
	NoRoute    bool
	Middleware bool
	Handlers   []gin.HandlerFunc
}
