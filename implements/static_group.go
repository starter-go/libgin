package implements

import "github.com/starter-go/libgin"

// StaticGroup 是默认的静态资源控制器组
type StaticGroup struct {
	//starter:component
	_as func(libgin.Group) //starter:as(".")

	Context libgin.Context //starter:inject("#")
	Path    string         //starter:inject("${web-group.static.path}")
	Name    string         //starter:inject("${web-group.static.name}")

	group group
}

// Registration ...
func (inst *StaticGroup) Registration() *libgin.GroupRegistration {
	return &libgin.GroupRegistration{
		Name:  inst.Name,
		Path:  inst.Path,
		Group: inst,
	}
}

// Route ...
func (inst *StaticGroup) Route(r libgin.Router) error {
	inst.group.context = inst.Context
	inst.group.path = inst.Path
	inst.group.name = inst.Name
	return inst.group.route(r)
}
