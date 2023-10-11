package controllers

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/application"
	"github.com/starter-go/base/util"
	"github.com/starter-go/libgin"
	"github.com/starter-go/libgin/devtools/views"
	"github.com/starter-go/libgin/devtools/xos"
)

// const 开发工具分组 = "devtools"

// EnvironmentController ...
type EnvironmentController struct {
	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Context   application.Context //starter:inject("context")
	Responder libgin.Responder    //starter:inject("#")
}

func (inst *EnvironmentController) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *EnvironmentController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Groups: []string{开发工具分组},
		Route:  inst.配置路由,
	}
}

func (inst *EnvironmentController) 配置路由(g libgin.RouterProxy) error {
	g = g.For("env")
	g.POST("", inst.handlePost)
	g.GET("", inst.handleGetList)
	g.GET(":id", inst.handleGetOne)
	g.PUT(":id", inst.handlePut)
	g.DELETE(":id", inst.handleDelete)
	return nil
}

func (inst *EnvironmentController) 执行(req *myEnvironmentRequest, fn func() error) {
	err := req.open()
	if err == nil {
		err = fn()
	}
	req.send(err)
}

func (inst *EnvironmentController) handlePost(c *gin.Context) {
	req := &myEnvironmentRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	inst.执行(req, req.todo)
}

func (inst *EnvironmentController) handlePut(c *gin.Context) {
	req := &myEnvironmentRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: true,
	}
	inst.执行(req, req.todo)
}

func (inst *EnvironmentController) handleDelete(c *gin.Context) {
	req := &myEnvironmentRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: false,
	}
	inst.执行(req, req.todo)
}

func (inst *EnvironmentController) handleGetList(c *gin.Context) {
	req := &myEnvironmentRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	inst.执行(req, req.doGetList)
}

func (inst *EnvironmentController) handleGetOne(c *gin.Context) {
	req := &myEnvironmentRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: false,
	}
	inst.执行(req, req.todo)
}

////////////////////////////////////////////////////////////////////////////////

type myEnvironmentRequest struct {
	context    *gin.Context
	controller *EnvironmentController

	wantRequestID   bool
	wantRequestBody bool

	id    xos.EnvID
	body1 views.Environment
	body2 views.Environment
}

func (inst *myEnvironmentRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		n, err := strconv.Atoi(idstr)
		if err != nil {
			return err
		}
		inst.id = xos.EnvID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body2)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myEnvironmentRequest) send(err error) {
	data := &inst.body2
	res := &libgin.Response{
		Context: inst.context,
		Data:    data,
		Error:   err,
		Status:  0,
	}
	inst.controller.Responder.Send(res)
}

func (inst *myEnvironmentRequest) todo() error {
	return nil
}

func (inst *myEnvironmentRequest) doGetList() error {

	env := inst.controller.Context.GetEnvironment()
	src := env.Export(nil)
	dst := inst.body2.Items

	for k, v := range src {
		dst = append(dst, &xos.Env{Name: k, Value: v})
	}

	sorter := &util.Sorter{}
	sorter.OnLen = func() int { return len(dst) }
	sorter.OnSwap = func(i1, i2 int) { dst[i1], dst[i2] = dst[i2], dst[i1] }
	sorter.OnLess = func(i1, i2 int) bool {
		k1 := dst[i1].Name
		k2 := dst[i2].Name
		return strings.Compare(k1, k2) < 0
	}
	sorter.Sort()

	inst.body2.Items = dst
	return nil
}
