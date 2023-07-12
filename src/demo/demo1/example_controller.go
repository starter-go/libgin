package demo1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/libgin/web"
	"github.com/starter-go/rest"
)

// ExampleController ...
type ExampleController struct {
	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")

}

func (inst *ExampleController) _impl() { inst._as(inst) }

// Registration ...
func (inst *ExampleController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *ExampleController) route(g *gin.RouterGroup) error {
	g = g.Group("example1")
	g.POST("", inst.handlePost)
	g.GET("", inst.handleGetList)
	g.GET(":id", inst.handleGetOne)
	g.PUT(":id", inst.handlePut)
	g.DELETE(":id", inst.handleDelete)
	return nil
}

func (inst *ExampleController) handleGetOne(c *gin.Context) {
	req := &myExampleRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	err := req.open()
	if err == nil {
		err = req.todo()
	}
	req.send(err)
}

func (inst *ExampleController) handleGetList(c *gin.Context) {
	req := &myExampleRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	err := req.open()
	if err == nil {
		err = req.todo()
	}
	req.send(err)
}

func (inst *ExampleController) handlePost(c *gin.Context) {
}

func (inst *ExampleController) handlePut(c *gin.Context) {
}

func (inst *ExampleController) handleDelete(c *gin.Context) {
}

////////////////////////////////////////////////////////////////////////////////

// ExampleVO ...
type ExampleVO struct {
	rest.BaseVO

	Items []*ExampleDTO `json:"items"`
}

// ExampleDTO ...
type ExampleDTO struct {
	rest.BaseDTO
}

////////////////////////////////////////////////////////////////////////////////

type myExampleRequest struct {
	context    *gin.Context
	controller *ExampleController

	wantRequestID   bool
	wantRequestBody bool

	id    rest.UserID
	body1 ExampleVO
	body2 ExampleVO
}

func (inst *myExampleRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		n, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = rest.UserID(n)
	}

	if inst.wantRequestBody {
		err := c.Bind(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myExampleRequest) send(err error) {
	ctx := inst.context
	status := inst.body2.Status
	data := &inst.body2
	resp := &libgin.Response{
		Context: ctx,
		Data:    data,
		Error:   err,
		Status:  status,
	}
	inst.controller.Responder.Send(resp)
}

func (inst *myExampleRequest) todo() error {
	id := inst.id
	if id > 1000 {
		return fmt.Errorf("err.id=%v", id)
	} else if id > 100 {
		return web.NewError(http.StatusForbidden, "err.id=%v", id)
	}
	return nil
}
