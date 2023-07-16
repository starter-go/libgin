package libgin

// ContentTypeRegistration 是文档类型信息
type ContentTypeRegistration struct {
	ContentType string
	Suffixes    []string
}

// ContentTypeRegistry 是文档类型登记接口
type ContentTypeRegistry interface {
	ListRegistrations() []*ContentTypeRegistration
}

// ContentTypeManager 是文档类型管理器
type ContentTypeManager interface {
	FindTypeBySuffix(suffix string) (string, error)
}
