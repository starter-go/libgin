package libgin

// Controller 表示一个 gin 控制器
type Controller interface {
	Registration() *ControllerRegistration
}

// ControllerRegistration 是控制器注册信息
type ControllerRegistration struct {
	Groups []string                   // 分组名称, 如果为空表示默认分组
	Route  func(rp RouterProxy) error // 控制器
}
