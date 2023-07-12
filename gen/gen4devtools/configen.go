//starter:configen(version="4")

package gen4devtools

import "github.com/starter-go/application"

// ExportForDevtools ...
func ExportForDevtools(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
