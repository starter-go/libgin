package implements

import (
	"fmt"

	"github.com/starter-go/application"
	"github.com/starter-go/libgin"
)

// DefaultContext ...
type DefaultContext struct {
	//starter:component
	_as func(libgin.Context) //starter:as("#")

	Controllers      []libgin.Controller        //starter:inject(".")
	Connectors       []libgin.ConnectorRegistry //starter:inject(".")
	Routers          []libgin.RouterRegistry    //starter:inject(".")
	Groups           []libgin.GroupRegistry     //starter:inject(".")
	DefaultGroupName string                     //starter:inject("${web.default-group-name}")

	inner *contextInner
}

func (inst *DefaultContext) _impl() (application.Lifecycle, libgin.Context) {
	return inst, inst
}

// Life ...
func (inst *DefaultContext) Life() *application.Life {
	return &application.Life{
		OnCreate: inst.onCreate,
	}
}

func (inst *DefaultContext) onCreate() error {
	inst.getInner()
	return nil
}

func (inst *DefaultContext) getInner() *contextInner {
	inner := inst.inner
	if inner == nil {
		i, err := inst.loadInner()
		if err != nil {
			panic(err)
		}
		inner = i
		inst.inner = inner
	}
	return inner
}

func (inst *DefaultContext) loadInner() (*contextInner, error) {

	ctx := &contextInner{}
	ctx.init()

	err := ctx.loadConnectors(inst.Connectors)
	if err != nil {
		return nil, err
	}

	err = ctx.loadGroups(inst.Groups)
	if err != nil {
		return nil, err
	}

	err = ctx.loadControllers(inst.Controllers)
	if err != nil {
		return nil, err
	}

	err = ctx.loadRouters(inst.Routers)
	if err != nil {
		return nil, err
	}

	return ctx, nil
}

// GetConnectorByName ...
func (inst *DefaultContext) GetConnectorByName(name string) (libgin.Connector, error) {
	inner := inst.getInner()
	table := inner.connectors
	err := fmt.Errorf("no libgin.Connector named '%s'", name)
	item := table[name]
	if item == nil {
		return nil, err
	}
	conn := item.Connector
	if conn == nil {
		return nil, err
	}
	return conn, nil
}

// GetGroupByName ...
func (inst *DefaultContext) GetGroupByName(name string) (libgin.Group, error) {
	inner := inst.getInner()
	table := inner.groups
	err := fmt.Errorf("no libgin.Group named '%s'", name)
	item := table[name]
	if item == nil {
		return nil, err
	}
	g := item.Group
	if g == nil {
		return nil, err
	}
	return g, nil
}

// GetRouterByName ...
func (inst *DefaultContext) GetRouterByName(name string) (libgin.Router, error) {
	inner := inst.getInner()
	table := inner.routers
	err := fmt.Errorf("no libgin.Router named '%s'", name)
	item := table[name]
	if item == nil {
		return nil, err
	}
	r := item.Router
	if r == nil {
		return nil, err
	}
	return r, nil
}

// ListControllersForGroup ...
func (inst *DefaultContext) ListControllersForGroup(wantGroupName string) ([]*libgin.ControllerRegistration, error) {
	inner := inst.getInner()
	list1 := inner.controllers
	list2 := make([]*libgin.ControllerRegistration, 0)
	for _, item := range list1 {
		if inst.isGroupNameMatched(wantGroupName, item.Groups) {
			list2 = append(list2, item)
		}
	}
	return list2, nil
}

func (inst *DefaultContext) isGroupNameMatched(wantGroupName string, havelist []string) bool {

	defaultGroupName := inst.DefaultGroupName

	if len(havelist) == 0 {
		havelist = []string{defaultGroupName}
	}

	for _, haveGroupName := range havelist {
		if haveGroupName == "" {
			haveGroupName = defaultGroupName
		}
		if wantGroupName == haveGroupName {
			return true
		}
	}

	return false
}

////////////////////////////////////////////////////////////////////////////////

// contextInner ...
type contextInner struct {
	controllers []*libgin.ControllerRegistration
	connectors  map[string]*libgin.ConnectorRegistration
	routers     map[string]*libgin.RouterRegistration
	groups      map[string]*libgin.GroupRegistration
}

func (inst *contextInner) init() {
	inst.controllers = make([]*libgin.ControllerRegistration, 0)
	inst.connectors = make(map[string]*libgin.ConnectorRegistration)
	inst.routers = make(map[string]*libgin.RouterRegistration)
	inst.groups = make(map[string]*libgin.GroupRegistration)
}

func (inst *contextInner) loadControllers(src []libgin.Controller) error {
	dst := inst.controllers
	for _, r1 := range src {
		r2 := r1.Registration()
		dst = append(dst, r2)
	}
	inst.controllers = dst
	return nil
}

func (inst *contextInner) loadConnectors(src []libgin.ConnectorRegistry) error {
	dst := inst.connectors
	r2list := make([]*libgin.ConnectorRegistration, 0)
	for _, r1 := range src {
		list := r1.ListRegistrations()
		r2list = append(r2list, list...)
	}
	for _, r2 := range r2list {
		name := r2.Name
		older := dst[name]
		if older != nil {
			return fmt.Errorf("the name of libgin.Connector is duplicate: %s", name)
		}
		dst[name] = r2
	}
	return nil
}

func (inst *contextInner) loadRouters(src []libgin.RouterRegistry) error {
	dst := inst.routers
	r2list := make([]*libgin.RouterRegistration, 0)
	for _, r1 := range src {
		list := r1.ListRegistrations()
		r2list = append(r2list, list...)
	}
	for _, r2 := range r2list {
		name := r2.Name
		older := dst[name]
		if older != nil {
			return fmt.Errorf("the name of libgin.Router is duplicate: %s", name)
		}
		dst[name] = r2
	}
	return nil
}

func (inst *contextInner) loadGroups(src []libgin.GroupRegistry) error {
	dst := inst.groups
	r2list := make([]*libgin.GroupRegistration, 0)
	for _, r1 := range src {
		list := r1.ListRegistrations()
		r2list = append(r2list, list...)
	}
	for _, r2 := range r2list {
		name := r2.Name
		older := dst[name]
		if older != nil {
			return fmt.Errorf("the name of libgin.Group is duplicate: %s", name)
		}
		dst[name] = r2
	}
	return nil
}
