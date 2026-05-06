package demo1

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/base/context2"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"
)

// ContextBindingDemo ...
type ContextBindingDemo struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")

}

func (inst *ContextBindingDemo) _impl() { inst._as(inst) }

// Registration ...
func (inst *ContextBindingDemo) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *ContextBindingDemo) route(g libgin.RouterProxy) error {

	g = g.For("demo/context-binding")

	g.GET("", inst.handleGet)

	// g.POST("", inst.handlePost)
	// g.GET(":id", inst.handleGetOne)
	// g.PUT(":id", inst.handlePut)
	// g.DELETE(":id", inst.handleDelete)

	return nil
}

func (inst *ContextBindingDemo) handleGet(c *gin.Context) {

	req := &myContextBindingDemoRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}

	req.execute(req.tryDemo)

}

////////////////////////////////////////////////////////////////////////////////

// ExampleVO ...
type ContextBindingDemoVO struct {

	// base
	rbac.BaseVO

	Items []string `json:"items"`
}

////////////////////////////////////////////////////////////////////////////////

type myContextBindingDemoRequest struct {
	context    *gin.Context
	controller *ContextBindingDemo

	wantRequestID   bool
	wantRequestBody bool

	id    int64
	body1 ContextBindingDemoVO
	body2 ContextBindingDemoVO
}

func (inst *myContextBindingDemoRequest) execute(task func() error) {

	ex := new(libgin.Executor)
	ex.Context = inst.context
	ex.Responder = inst.controller.Responder
	ex.Body1 = &inst.body1
	ex.Body2 = &inst.body2

	ex.OnTask = task
	ex.OnOpen = inst.onOpen

	ex.Execute()
}

func (inst *myContextBindingDemoRequest) onOpen(c *gin.Context) error {

	// c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		n, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = n
	}

	if inst.wantRequestBody {
		err := c.Bind(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myContextBindingDemoRequest) tryDemo() error {

	var cc context.Context
	cc = inst.context

	values, err := context2.GetValues(cc)
	if err != nil {
		return err
	}

	gc, err := libgin.GetGinContext(cc)
	if err != nil {
		return err
	}

	values.SetValue("attr(float)", 3.14159)
	values.SetValue("attr(string)", "hello,context2")
	values.SetValue("attr([]byte)", []byte{2, 3, 5, 7, 11, 13})

	keys := values.Keys()
	url := gc.Request.URL
	strlist := inst.body2.Items

	strlist = append(strlist, "url = "+url.String())
	strlist = append(strlist, "keys:")
	strlist = append(strlist, keys...)

	inst.body2.Items = strlist
	return nil
}
