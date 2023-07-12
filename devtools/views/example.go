package views

import "github.com/starter-go/libgin/devtools/xos"

// Example 是 Example 的 VO
type Example struct {
	Base

	Items []*xos.Example `json:"items"`
}
