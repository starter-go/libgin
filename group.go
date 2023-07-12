package libgin

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
