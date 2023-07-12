package views

import "github.com/starter-go/libgin/devtools/xos"

// Components 是 Components 的 VO
type Components struct {
	Base

	Items []*xos.Component `json:"components"`
}
