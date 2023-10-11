package implements

import (
	"github.com/starter-go/libgin"
)

type group struct {
	context libgin.Context
	path    string
	name    string
}

func (inst *group) route(r libgin.Router) error {

	proxy, ctx, err := inst.prepareRC(r)
	if err != nil {
		return err
	}

	crlist, err := inst.listControllers()
	if err != nil {
		return err
	}

	for _, cr := range crlist {
		fn := cr.Route
		if fn == nil {
			continue
		}
		err = fn(proxy)
		if err != nil {
			return err
		}
	}

	return ctx.apply()
}

func (inst *group) prepareRC(r libgin.Router) (libgin.RouterProxy, *routerContext, error) {
	engine := r.Engine()
	proxy1 := makeProxyForRouter(engine)
	proxy2 := proxy1.For(inst.path)
	return proxy2, proxy1.context, nil
}

// func (inst *group) prepareGroup(r libgin.Router) (*gin.RouterGroup, error) {
// 	engine := r.Engine()
// 	g := engine.Group(inst.path)
// 	return g, nil
// }

func (inst *group) listControllers() ([]*libgin.ControllerRegistration, error) {
	return inst.context.ListControllersForGroup(inst.name)
}

////////////////////////////////////////////////////////////////////////////////

type commonGroup struct {
	group group
}

func (inst *commonGroup) Registration() *libgin.GroupRegistration {
	return &libgin.GroupRegistration{
		Name:  inst.group.name,
		Path:  inst.group.path,
		Group: inst,
	}
}

func (inst *commonGroup) Route(r libgin.Router) error {
	return inst.group.route(r)
}

////////////////////////////////////////////////////////////////////////////////

// GroupBuilder 用于创建 libgin.Group 实例
type GroupBuilder struct {
	Context libgin.Context
	Name    string
	Path    string
}

// Create 创建 libgin.Group 实例
func (inst *GroupBuilder) Create() libgin.Group {
	g := &commonGroup{}
	g.group.context = inst.Context
	g.group.name = inst.Name
	g.group.path = inst.Path
	return g
}
