//starter:configen(version="4")

package gen4demo

import "github.com/starter-go/application"

// ExportForDemo ...
func ExportForDemo(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
