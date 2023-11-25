package groups

import (
	"github.com/starter-go/libgin"
	"github.com/starter-go/libgin/implements"
)

// DevtoolsGroup ...
type DevtoolsGroup struct {

	//starter:component
	_as func(libgin.Group) //starter:as(".")

	Context libgin.Context //starter:inject("#")
	Enabled bool           //starter:inject("${web-group.devtools.enabled}")
}

// Registration ...
func (inst *DevtoolsGroup) Registration() *libgin.GroupRegistration {
	return &libgin.GroupRegistration{
		Name:    "devtools",
		Path:    "/api/devtools/",
		Group:   inst,
		Enabled: inst.Enabled,
	}
}

// Route ...
func (inst *DevtoolsGroup) Route(r libgin.Router) error {

	reg := inst.Registration()
	gb := &implements.GroupBuilder{}
	gb.Name = reg.Name
	gb.Path = reg.Path
	gb.Context = inst.Context

	g := gb.Create()
	return g.Route(r)
}
