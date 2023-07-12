package libgin

// Context 是 libgin 上下文，它集中管理各种 libgin 资源
type Context interface {
	GetConnectorByName(name string) (Connector, error)
	GetGroupByName(name string) (Group, error)
	GetRouterByName(name string) (Router, error)
	ListControllersForGroup(groupName string) ([]*ControllerRegistration, error)
}
