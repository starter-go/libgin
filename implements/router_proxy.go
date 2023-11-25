package implements

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

func makeProxyForRouter(engine *gin.Engine) *routerProxy {
	ctx := &routerContext{
		engine: engine,
	}
	proxy := &routerProxy{
		context: ctx,
		path:    "/",
	}
	return proxy
}

////////////////////////////////////////////////////////////////////////////////

type routerContext struct {
	engine      *gin.Engine
	routingList []*libgin.Routing
}

func (inst *routerContext) apply() error {
	sort.Sort(inst)
	list := inst.routingList
	for _, r := range list {
		if r.Middleware {
			inst.engine.Use(r.Handlers...)
		} else if r.NoMethod {
			inst.engine.NoMethod(r.Handlers...)
		} else if r.NoRoute {
			inst.engine.NoRoute(r.Handlers...)
		} else {
			inst.engine.Handle(r.Method, r.Path, r.Handlers...)
		}
	}
	return nil
}

func (inst *routerContext) Swap(i1, i2 int) {
	list := inst.routingList
	list[i1], list[i2] = list[i2], list[i1]
}

func (inst *routerContext) Len() int {
	return len(inst.routingList)
}

func (inst *routerContext) Less(i1, i2 int) bool {

	list := inst.routingList
	o1 := list[i1]
	o2 := list[i2]
	p1 := o1.Priority
	p2 := o2.Priority

	if p1 < p2 {
		return false
	} else if p1 > p2 {
		return true
	}

	s1 := o1.Path + "#" + o1.Method
	s2 := o2.Path + "#" + o2.Method
	return strings.Compare(s1, s2) < 0
}

////////////////////////////////////////////////////////////////////////////////

type routerProxy struct {
	context *routerContext
	path    string
}

func (inst *routerProxy) _impl() libgin.RouterProxy {
	return inst
}

func (inst *routerProxy) normalizePath(path string) string {
	const sep = "/"
	if strings.HasPrefix(path, sep) {
		// NOP
	} else {
		path = inst.path + sep + path
	}
	parts := strings.Split(path, sep)
	builder := strings.Builder{}
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		builder.WriteString(sep)
		builder.WriteString(part)
	}
	return builder.String()
}

func (inst *routerProxy) For(path string) libgin.RouterProxy {
	path2 := inst.normalizePath(path)
	return &routerProxy{
		context: inst.context,
		path:    path2,
	}
}

func (inst *routerProxy) Route(r *libgin.Routing) {
	list := inst.context.routingList
	path2 := inst.normalizePath(r.Path)
	r.Path = path2
	list = append(list, r)
	if r.AsIndexPage {
		r2, err := inst.makeRoutingForIndexPage(r)
		if err == nil {
			list = append(list, r2)
		}
	}
	inst.context.routingList = list
}

func (inst *routerProxy) makeRoutingForIndexPage(src *libgin.Routing) (*libgin.Routing, error) {

	path := src.Path
	i := strings.LastIndex(path, "/")
	if i < 0 {
		return nil, fmt.Errorf("bad index page path: %s", path)
	}
	path = path[0 : i+1]

	dst := &libgin.Routing{}
	*dst = *src
	dst.Path = path
	return dst, nil
}

func (inst *routerProxy) Handle(method string, path string, h ...gin.HandlerFunc) {
	r := &libgin.Routing{
		Method:   method,
		Path:     path,
		Handlers: h,
	}
	inst.Route(r)
}

func (inst *routerProxy) NoMethod(h ...gin.HandlerFunc) {
	r := &libgin.Routing{
		NoMethod: true,
		Handlers: h,
	}
	inst.Route(r)
}

func (inst *routerProxy) NoRoute(h ...gin.HandlerFunc) {
	r := &libgin.Routing{
		NoRoute:  true,
		Handlers: h,
	}
	inst.Route(r)
}

func (inst *routerProxy) GET(path string, h ...gin.HandlerFunc) {
	r := &libgin.Routing{
		Method:   http.MethodGet,
		Path:     path,
		Handlers: h,
	}
	inst.Route(r)
}

func (inst *routerProxy) POST(path string, h ...gin.HandlerFunc) {
	r := &libgin.Routing{
		Method:   http.MethodPost,
		Path:     path,
		Handlers: h,
	}
	inst.Route(r)
}

func (inst *routerProxy) PUT(path string, h ...gin.HandlerFunc) {
	r := &libgin.Routing{
		Method:   http.MethodPut,
		Path:     path,
		Handlers: h,
	}
	inst.Route(r)
}

func (inst *routerProxy) DELETE(path string, h ...gin.HandlerFunc) {
	r := &libgin.Routing{
		Method:   http.MethodDelete,
		Path:     path,
		Handlers: h,
	}
	inst.Route(r)
}

////////////////////////////////////////////////////////////////////////////////
