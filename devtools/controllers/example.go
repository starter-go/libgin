package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/application"
	"github.com/starter-go/libgin"
	"github.com/starter-go/libgin/devtools/views"
	"github.com/starter-go/libgin/devtools/xos"
)

// const 开发工具分组 = "devtools"

// ExampleController ...
type ExampleController struct {
	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Context   application.Context //starter:inject("context")
	Responder libgin.Responder    //starter:inject("#")
}

func (inst *ExampleController) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *ExampleController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Groups: []string{开发工具分组},
		Route:  inst.配置路由,
	}
}

func (inst *ExampleController) 配置路由(g libgin.RouterProxy) error {
	g = g.For("example")
	g.POST("", inst.handlePost)
	g.GET("", inst.handleGetList)
	g.GET(":id", inst.handleGetOne)
	g.PUT(":id", inst.handlePut)
	g.DELETE(":id", inst.handleDelete)
	return nil
}

func (inst *ExampleController) 执行(req *myExampleRequest, fn func() error) {
	err := req.open()
	if err == nil {
		err = fn()
	}
	req.send(err)
}

func (inst *ExampleController) handlePost(c *gin.Context) {
	req := &myExampleRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	inst.执行(req, req.todo)
}

func (inst *ExampleController) handlePut(c *gin.Context) {
	req := &myExampleRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: true,
	}
	inst.执行(req, req.todo)
}

func (inst *ExampleController) handleDelete(c *gin.Context) {
	req := &myExampleRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: false,
	}
	inst.执行(req, req.todo)
}

func (inst *ExampleController) handleGetList(c *gin.Context) {
	req := &myExampleRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	inst.执行(req, req.todo)
}

func (inst *ExampleController) handleGetOne(c *gin.Context) {
	req := &myExampleRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: false,
	}
	inst.执行(req, req.todo)
}

////////////////////////////////////////////////////////////////////////////////

type myExampleRequest struct {
	context    *gin.Context
	controller *ExampleController

	wantRequestID   bool
	wantRequestBody bool

	id    xos.ExampleID
	body1 views.Example
	body2 views.Example
}

func (inst *myExampleRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		n, err := strconv.Atoi(idstr)
		if err != nil {
			return err
		}
		inst.id = xos.ExampleID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body2)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myExampleRequest) send(err error) {
	data := &inst.body2
	res := &libgin.Response{
		Context: inst.context,
		Data:    data,
		Error:   err,
		Status:  0,
	}
	inst.controller.Responder.Send(res)
}

func (inst *myExampleRequest) todo() error {
	return nil
}
