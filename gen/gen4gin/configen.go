//starter:configen(version="4")

package gen4gin

import "github.com/starter-go/application"

// ExportForGin ...
func ExportForGin(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
