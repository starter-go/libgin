package xos

import "github.com/starter-go/application/components"

// ComponentID ...
type ComponentID int

// Component ...
type Component struct {
	Base

	ID      components.ID `json:"id"`
	Scope   string        `json:"scope"`
	Classes []string      `json:"classes"`
	Aliases []string      `json:"aliases"`
}
