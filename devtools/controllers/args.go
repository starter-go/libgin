package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/application"
	"github.com/starter-go/libgin"
	"github.com/starter-go/libgin/devtools/views"
	"github.com/starter-go/libgin/devtools/xos"
)

const 开发工具分组 = "devtools"

// ArgumentController ...
type ArgumentController struct {
	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Context   application.Context //starter:inject("context")
	Responder libgin.Responder    //starter:inject("#")
}

func (inst *ArgumentController) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *ArgumentController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Groups: []string{开发工具分组},
		Route:  inst.配置路由,
	}
}

func (inst *ArgumentController) 配置路由(g libgin.RouterProxy) error {
	g = g.For("arguments")
	g.POST("", inst.handlePost)
	g.GET("", inst.handleGetList)
	g.GET(":id", inst.handleGetOne)
	g.PUT(":id", inst.handlePut)
	g.DELETE(":id", inst.handleDelete)
	return nil
}

func (inst *ArgumentController) 执行(req *myArgumentRequest, fn func() error) {
	err := req.open()
	if err == nil {
		err = fn()
	}
	req.send(err)
}

func (inst *ArgumentController) handlePost(c *gin.Context) {
	req := &myArgumentRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	inst.执行(req, req.todo)
}

func (inst *ArgumentController) handlePut(c *gin.Context) {
	req := &myArgumentRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: true,
	}
	inst.执行(req, req.todo)
}

func (inst *ArgumentController) handleDelete(c *gin.Context) {
	req := &myArgumentRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: false,
	}
	inst.执行(req, req.todo)
}

func (inst *ArgumentController) handleGetList(c *gin.Context) {
	req := &myArgumentRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	inst.执行(req, req.doGetList)
}

func (inst *ArgumentController) handleGetOne(c *gin.Context) {
	req := &myArgumentRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: false,
	}
	inst.执行(req, req.todo)
}

////////////////////////////////////////////////////////////////////////////////

type myArgumentRequest struct {
	context    *gin.Context
	controller *ArgumentController

	wantRequestID   bool
	wantRequestBody bool

	id    xos.ArgumentID
	body1 views.Arguments
	body2 views.Arguments
}

func (inst *myArgumentRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		n, err := strconv.Atoi(idstr)
		if err != nil {
			return err
		}
		inst.id = xos.ArgumentID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body2)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myArgumentRequest) send(err error) {
	data := &inst.body2
	res := &libgin.Response{
		Context: inst.context,
		Data:    data,
		Error:   err,
		Status:  0,
	}
	inst.controller.Responder.Send(res)
}

func (inst *myArgumentRequest) todo() error {
	return nil
}

func (inst *myArgumentRequest) doGetList() error {
	args := inst.controller.Context.GetArguments()
	inst.body2.Args = args.Raw()
	return nil
}
