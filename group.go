package libgin

// 定义常用的分组名称
const (
	GroupREST   = "rest"
	GroupStatic = "static"
)

// Group ...
type Group interface {
	// Registration( ... ) *GroupRegistration
	Route(r Router) error
}

// GroupRegistry ...
type GroupRegistry interface {
	ListRegistrations() []*GroupRegistration
}

// GroupRegistration ...
type GroupRegistration struct {
	Enabled bool
	Name    string
	Path    string
	Group   Group
}
