package libgin

import "github.com/gin-gonic/gin"

// Router ...
type Router interface {
	// Registration(  ) *RouterRegistration
	Engine() *gin.Engine
}

// RouterRegistry ...
type RouterRegistry interface {
	ListRegistrations() []*RouterRegistration
}

// RouterRegistration ...
type RouterRegistration struct {
	Name   string
	Router Router
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
	Method      string
	Path        string
	Priority    int
	NoMethod    bool
	NoRoute     bool
	Middleware  bool
	AsIndexPage bool
	Handlers    []gin.HandlerFunc
}
