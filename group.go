package libgin

// 定义常用的分组名称
const (
	GroupREST   = "rest"
	GroupStatic = "static"
)

// Group ...
type Group interface {
	Registration() *GroupRegistration
	Route(r Router) error
}

// GroupRegistration ...
type GroupRegistration struct {
	Name  string
	Path  string
	Group Group
}
