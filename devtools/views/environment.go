package views

import "github.com/starter-go/libgin/devtools/xos"

// Environment 是 Environment 的 VO
type Environment struct {
	Base

	Items []*xos.Env `json:"env"`
}
