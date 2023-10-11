package controllers

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/application"
	"github.com/starter-go/application/components"
	"github.com/starter-go/base/util"
	"github.com/starter-go/libgin"
	"github.com/starter-go/libgin/devtools/views"
	"github.com/starter-go/libgin/devtools/xos"
)

// const 开发工具分组 = "devtools"

// ComponentController ...
type ComponentController struct {
	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Context   application.Context //starter:inject("context")
	Responder libgin.Responder    //starter:inject("#")
}

func (inst *ComponentController) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *ComponentController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Groups: []string{开发工具分组},
		Route:  inst.配置路由,
	}
}

func (inst *ComponentController) 配置路由(g libgin.RouterProxy) error {
	g = g.For("components")
	g.POST("", inst.handlePost)
	g.GET("", inst.handleGetList)
	g.GET(":id", inst.handleGetOne)
	g.PUT(":id", inst.handlePut)
	g.DELETE(":id", inst.handleDelete)
	return nil
}

func (inst *ComponentController) 执行(req *myComponentRequest, fn func() error) {
	err := req.open()
	if err == nil {
		err = fn()
	}
	req.send(err)
}

func (inst *ComponentController) handlePost(c *gin.Context) {
	req := &myComponentRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	inst.执行(req, req.todo)
}

func (inst *ComponentController) handlePut(c *gin.Context) {
	req := &myComponentRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: true,
	}
	inst.执行(req, req.todo)
}

func (inst *ComponentController) handleDelete(c *gin.Context) {
	req := &myComponentRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: false,
	}
	inst.执行(req, req.todo)
}

func (inst *ComponentController) handleGetList(c *gin.Context) {
	req := &myComponentRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	inst.执行(req, req.doGetList)
}

func (inst *ComponentController) handleGetOne(c *gin.Context) {
	req := &myComponentRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: false,
	}
	inst.执行(req, req.todo)
}

////////////////////////////////////////////////////////////////////////////////

type myComponentRequest struct {
	context    *gin.Context
	controller *ComponentController

	wantRequestID   bool
	wantRequestBody bool

	id    xos.ComponentID
	body1 views.Components
	body2 views.Components
}

func (inst *myComponentRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		n, err := strconv.Atoi(idstr)
		if err != nil {
			return err
		}
		inst.id = xos.ComponentID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body2)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myComponentRequest) send(err error) {
	data := &inst.body2
	res := &libgin.Response{
		Context: inst.context,
		Data:    data,
		Error:   err,
		Status:  0,
	}
	inst.controller.Responder.Send(res)
}

func (inst *myComponentRequest) todo() error {
	return nil
}

func (inst *myComponentRequest) doGetList() error {

	table := inst.controller.Context.GetComponents()
	src := table.Export(nil)
	dst := inst.body2.Items

	for _, h := range src {
		info := h.Info()
		item := &xos.Component{}
		item.ID = info.ID()
		item.Aliases = info.Aliases()
		item.Classes = info.Classes()
		item.Scope = inst.stringifyScope(info.Scope())
		dst = append(dst, item)
	}

	sorter := &util.Sorter{}
	sorter.OnLen = func() int { return len(dst) }
	sorter.OnSwap = func(i1, i2 int) { dst[i1], dst[i2] = dst[i2], dst[i1] }
	sorter.OnLess = func(i1, i2 int) bool {
		k1 := dst[i1].ID.String()
		k2 := dst[i2].ID.String()
		return strings.Compare(k1, k2) < 0
	}
	sorter.Sort()

	inst.body2.Items = dst
	return nil
}

func (inst *myComponentRequest) stringifyScope(scope components.Scope) string {
	switch scope {
	case components.ScopeSingleton:
		return "Singleton"
	case components.ScopePrototype:
		return "Prototype"
	}
	return ""
}
