package libgin

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgin"
	"github.com/starter-go/libgin/gen/gen4devtools"
	"github.com/starter-go/libgin/gen/gen4gin"
)

// Module ...
func Module() application.Module {
	mb := libgin.NewModule()
	mb.Components(gen4gin.ExportForGin)
	return mb.Create()
}

// ModuleDevtools ...
func ModuleDevtools() application.Module {
	mb := libgin.NewModuleDevtools()
	mb.Components(gen4devtools.ExportForDevtools)
	mb.Depend(Module())
	return mb.Create()
}
