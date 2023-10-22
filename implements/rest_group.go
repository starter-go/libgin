package implements

import (
	"strings"

	"github.com/starter-go/application"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/libgin"
)

// RESTGroupRegistry 是默认的 RESTful-API 控制器组
type RESTGroupRegistry struct {

	//starter:component
	_as func(libgin.GroupRegistry) //starter:as(".")

	AC          application.Context //starter:inject("context")
	Context     libgin.Context      //starter:inject("#")
	GroupIDList string              //starter:inject("${web.groups}")

	// Name string //starter:inject("${web-group.rest.name}")
	// Path string //starter:inject("${web-group.rest.path}")

	cached []*libgin.GroupRegistration
}

func (inst *RESTGroupRegistry) _impl() libgin.GroupRegistry {
	return inst
}

// ListRegistrations ...
func (inst *RESTGroupRegistry) ListRegistrations() []*libgin.GroupRegistration {
	list := inst.cached
	if list == nil {
		newlist, err := inst.loadGroups()
		if err != nil {
			panic(err)
		}
		list = newlist
		inst.cached = newlist
	}
	return list
}

func (inst *RESTGroupRegistry) listGroupIDs() []string {
	str := inst.GroupIDList
	sep := ","
	list1 := strings.Split(str, sep)
	list2 := make([]string, 0)
	for _, id := range list1 {
		id = strings.TrimSpace(id)
		if id != "" {
			list2 = append(list2, id)
		}
	}
	return list2
}

func (inst *RESTGroupRegistry) loadGroups() ([]*libgin.GroupRegistration, error) {
	props := inst.AC.GetProperties()
	getter := props.Getter()
	r2list := make([]*libgin.GroupRegistration, 0)
	ids := inst.listGroupIDs()
	for _, id := range ids {
		r2, err := inst.loadGroup(id, getter.Required())
		if err != nil {
			return nil, err
		}
		r2list = append(r2list, r2)
	}
	return r2list, nil
}

func (inst *RESTGroupRegistry) loadGroup(id string, getter properties.Getter) (*libgin.GroupRegistration, error) {

	prefix := "web-group." + id + "."

	g := &restGroup{}
	g.context = inst.Context
	g.name = getter.GetString(prefix + "name")
	g.path = getter.GetString(prefix + "path")
	g.enabled = getter.GetBool(prefix + "enabled")

	err := getter.Error()
	if err != nil {
		return nil, err
	}
	r2 := &libgin.GroupRegistration{
		Enabled: g.enabled,
		Group:   g,
		Name:    g.name,
		Path:    g.path,
	}
	return r2, nil
}

////////////////////////////////////////////////////////////////////////////////

type restGroup struct {
	group
	enabled bool
}

// Route ...
func (inst *restGroup) Route(r libgin.Router) error {
	// inst.group.context = inst.Context
	// inst.group.name = inst.Name
	// inst.group.path = inst.Path
	return inst.group.route(r)
}
