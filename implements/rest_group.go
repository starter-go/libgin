package implements

import "github.com/starter-go/libgin"

// RESTGroup 是默认的 RESTful-API 控制器组
type RESTGroup struct {
	//starter:component
	_as func(libgin.Group) //starter:as(".")

	Context libgin.Context //starter:inject("#")
	Name    string         //starter:inject("${web-group.rest.name}")
	Path    string         //starter:inject("${web-group.rest.path}")

	group group
}

// Registration ...
func (inst *RESTGroup) Registration() *libgin.GroupRegistration {
	return &libgin.GroupRegistration{
		Name:  inst.Name,
		Path:  inst.Path,
		Group: inst,
	}
}

// Route ...
func (inst *RESTGroup) Route(r libgin.Router) error {
	inst.group.context = inst.Context
	inst.group.name = inst.Name
	inst.group.path = inst.Path
	return inst.group.route(r)
}
