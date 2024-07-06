package libgin

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgin"
	"github.com/starter-go/libgin/gen/demo4libgin"
	"github.com/starter-go/libgin/gen/devtools4libgin"
	"github.com/starter-go/libgin/gen/main4libgin"
	"github.com/starter-go/starter"
	"github.com/starter-go/stopper/modules/stopper"
)

// Module ...
func Module() application.Module {
	mb := libgin.NewMainModule()
	mb.Components(main4libgin.ExportComponents)
	mb.Depend(starter.Module())
	mb.Depend(stopper.Module())
	return mb.Create()
}

// ModuleDevtools ...
func ModuleDevtools() application.Module {
	mb := libgin.NewDevtoolsModule()
	mb.Components(devtools4libgin.ExportComponents)
	mb.Depend(Module())
	return mb.Create()
}

// ModuleDemo ...
func ModuleDemo() application.Module {
	mb := libgin.NewDemoModule()
	mb.Components(demo4libgin.ExportComponents)
	mb.Depend(Module())
	mb.Depend(ModuleDevtools())
	return mb.Create()
}
