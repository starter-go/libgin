package views

import "github.com/starter-go/libgin/devtools/xos"

// Properties 是 Properties 的 VO
type Properties struct {
	Base

	Items []*xos.Property `json:"properties"`
}
