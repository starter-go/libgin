package implements

import (
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/application"
	"github.com/starter-go/libgin"
	"github.com/starter-go/vlog"
)

// DefaultRouter 默认路由
type DefaultRouter struct {
	//starter:component
	_as func(libgin.Router) //starter:as(".")

	Context   libgin.Context //starter:inject("#")
	Name      string         //starter:inject("${web-router.default.name}")
	GroupList string         //starter:inject("${web-router.default.groups}")
	GinMode   string         //starter:inject("${gin.mode}")

	engine *gin.Engine
}

func (inst *DefaultRouter) _impl() (libgin.Router, application.Lifecycle) {
	return inst, inst
}

// Life ...
func (inst *DefaultRouter) Life() *application.Life {
	return &application.Life{
		OnCreate: inst.makeRoute,
	}
}

// Registration ...
func (inst *DefaultRouter) Registration() *libgin.RouterRegistration {
	return &libgin.RouterRegistration{
		Name:   inst.Name,
		Router: inst,
	}
}

func (inst *DefaultRouter) createEngine() *gin.Engine {
	mode := inst.GinMode
	switch mode {
	case gin.DebugMode:
		break
	case gin.ReleaseMode:
		break
	default:
		mode = gin.DebugMode
		vlog.Warn("the property: 'gin.mode' can be [debug|release]")
	}
	gin.SetMode(mode)
	return gin.Default()
}

// Engine ...
func (inst *DefaultRouter) Engine() *gin.Engine {
	e := inst.engine
	if e == nil {
		e = inst.createEngine()
		inst.engine = e
	}
	return e
}

func (inst *DefaultRouter) makeRoute() error {
	namelist := strings.Split(inst.GroupList, ",")
	sort.Strings(namelist)
	for _, name := range namelist {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}
		g, err := inst.Context.GetGroupByName(name)
		if err != nil {
			return err
		}
		err = g.Route(inst)
		if err != nil {
			return err
		}
	}
	return nil
}
