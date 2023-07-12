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

// PropertyController ...
type PropertyController struct {
	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Context   application.Context //starter:inject("context")
	Responder libgin.Responder    //starter:inject("#")
}

func (inst *PropertyController) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *PropertyController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Groups: []string{开发工具分组},
		Route:  inst.配置路由,
	}
}

func (inst *PropertyController) 配置路由(g *gin.RouterGroup) error {
	g = g.Group("properties")
	g.POST("", inst.handlePost)
	g.GET("", inst.handleGetList)
	g.GET(":id", inst.handleGetOne)
	g.PUT(":id", inst.handlePut)
	g.DELETE(":id", inst.handleDelete)
	return nil
}

func (inst *PropertyController) 执行(req *myPropertyRequest, fn func() error) {
	err := req.open()
	if err == nil {
		err = fn()
	}
	req.send(err)
}

func (inst *PropertyController) handlePost(c *gin.Context) {
	req := &myPropertyRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	inst.执行(req, req.todo)
}

func (inst *PropertyController) handlePut(c *gin.Context) {
	req := &myPropertyRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: true,
	}
	inst.执行(req, req.todo)
}

func (inst *PropertyController) handleDelete(c *gin.Context) {
	req := &myPropertyRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: false,
	}
	inst.执行(req, req.todo)
}

func (inst *PropertyController) handleGetList(c *gin.Context) {
	req := &myPropertyRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	inst.执行(req, req.doGetList)
}

func (inst *PropertyController) handleGetOne(c *gin.Context) {
	req := &myPropertyRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: false,
	}
	inst.执行(req, req.todo)
}

////////////////////////////////////////////////////////////////////////////////

type myPropertyRequest struct {
	context    *gin.Context
	controller *PropertyController

	wantRequestID   bool
	wantRequestBody bool

	id    xos.PropertyID
	body1 views.Properties
	body2 views.Properties
}

func (inst *myPropertyRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		n, err := strconv.Atoi(idstr)
		if err != nil {
			return err
		}
		inst.id = xos.PropertyID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body2)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myPropertyRequest) send(err error) {
	data := &inst.body2
	res := &libgin.Response{
		Context: inst.context,
		Data:    data,
		Error:   err,
		Status:  0,
	}
	inst.controller.Responder.Send(res)
}

func (inst *myPropertyRequest) todo() error {
	return nil
}

func (inst *myPropertyRequest) doGetList() error {

	ctx := inst.controller.Context
	src := ctx.GetProperties().Export(nil)
	dst := inst.body2.Items

	for k, v := range src {
		dst = append(dst, &xos.Property{
			Name:  k,
			Value: v,
		})
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
