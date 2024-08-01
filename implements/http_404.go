package implements

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/application"
	"github.com/starter-go/libgin"
	"github.com/starter-go/vlog"
)

// HTTP404Controller ...
type HTTP404Controller struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	AppContext application.Context //starter:inject("context")

	ResponseContentRes  string //starter:inject("${libgin.http404.page.resource}")
	ResponseContentType string //starter:inject("${libgin.http404.page.mediatype}")
	ResponseCode        int    //starter:inject("${libgin.http404.page.status}")
	HandlerPriority     int    //starter:inject("${libgin.http404.page.priority}")

	cache *my404ControllerCache
}

func (inst *HTTP404Controller) _impl() libgin.Controller {
	return inst
}

// Registration 。。。
func (inst *HTTP404Controller) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Route: inst.route,
	}
}

func (inst *HTTP404Controller) route(rp libgin.RouterProxy) error {

	hlist := []gin.HandlerFunc{inst.handle}
	priority := inst.HandlerPriority

	rp.Route(&libgin.Routing{
		NoRoute:  true,
		Priority: priority,
		Handlers: hlist,
	})

	rp.Route(&libgin.Routing{
		NoMethod: true,
		Priority: priority,
		Handlers: hlist,
	})

	return nil
}

func (inst *HTTP404Controller) handle(c *gin.Context) {
	x := inst.getCache()
	c.Data(x.code, x.mtype, x.data)
}

func (inst *HTTP404Controller) getCache() *my404ControllerCache {

	x := inst.cache
	if x != nil {
		return x
	}

	x, err := inst.loadCache()
	if err != nil {
		vlog.Error(err.Error())
		x = inst.loadCacheDefault()
	}

	inst.cache = x
	return x
}

func (inst *HTTP404Controller) loadCache() (*my404ControllerCache, error) {

	path := inst.ResponseContentRes
	code := inst.ResponseCode
	mtype := inst.ResponseContentType

	data, err := inst.AppContext.GetResources().ReadBinary(path)
	if err != nil {
		return nil, err
	}

	x := &my404ControllerCache{
		code:  code,
		mtype: mtype,
		data:  data,
	}
	return x, nil
}

func (inst *HTTP404Controller) loadCacheDefault() *my404ControllerCache {
	code := http.StatusNotFound
	msg := http.StatusText(code)
	content := fmt.Sprintf("HTTP %d %s", code, msg)
	return &my404ControllerCache{
		code:  code,
		mtype: "text/plain",
		data:  []byte(content),
	}
}

////////////////////////////////////////////////////////////////////////////////

type my404ControllerCache struct {
	data  []byte
	mtype string
	code  int
}
