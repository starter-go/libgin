package implements

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// GinContextBinderMiddleware 是一个中间件, 它基于 gin.Context, 提供对 context2.* 的支持
type GinContextBinderMiddleware struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Enabled  bool   //starter:inject("${web-middleware.context-binding.enabled}")
	Priority int    //starter:inject("${web-middleware.context-binding.priority}")
	Path     string //starter:inject("${web-middleware.context-binding.path}")

}

func (inst *GinContextBinderMiddleware) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *GinContextBinderMiddleware) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Route: inst.route,
	}
}

func (inst *GinContextBinderMiddleware) route(rp libgin.RouterProxy) error {

	if !inst.Enabled {
		return nil
	}

	r1 := &libgin.Routing{
		Middleware: true,
		Priority:   inst.Priority,
		Path:       inst.Path,
	}

	r1.Handlers = append(r1.Handlers, inst.handle)
	rp.Route(r1)
	return nil
}

func (inst *GinContextBinderMiddleware) handle(c *gin.Context) {

	_, err := libgin.SetupGinContext(c)

	if err != nil {
		const code = http.StatusInternalServerError
		c.AbortWithError(code, err)
	}

}
