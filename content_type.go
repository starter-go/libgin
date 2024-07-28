package libgin

// // MIMEType 结构包含关于一个类型的详细信息
// type MIMEType struct {
// 	Type        string // the type name , like 'image/png'
// 	Label       string
// 	Description string
// 	Icon        string
// }

// // ContentTypeRegistration 是文档类型信息
// type ContentTypeRegistration struct {
// 	ContentType string
// 	Info        *MIMEType
// 	Suffixes    []string
// }

// // ContentTypeRegistry 是文档类型登记接口
// type ContentTypeRegistry interface {
// 	ListRegistrations() []*ContentTypeRegistration
// }

// ContentTypeManager 是文档类型管理器
type ContentTypeManager interface {
	FindTypeBySuffix(suffix string) (string, error)
}
