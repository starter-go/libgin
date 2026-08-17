package libgin

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgin"
	"github.com/starter-go/libgin/gen/demo4libgin"
	"github.com/starter-go/libgin/gen/devtools4libgin"
	"github.com/starter-go/libgin/gen/main4libgin"
	"github.com/starter-go/libgin/gen/test4libgin"
	"github.com/starter-go/mimetypes/modules/mimetypes"
	"github.com/starter-go/starter"
	"github.com/starter-go/stopper/modules/stopper"
	"github.com/starter-go/v0/libvlog"
)

func Module() application.Module {
	return ModuleForMain()
}

func ModuleForMain() application.Module {
	mb := libgin.NewMainModule()
	mb.Components(main4libgin.ExportComponents)
	mb.Depend(starter.Module())
	mb.Depend(stopper.Module())
	mb.Depend(libvlog.Module())
	mb.Depend(mimetypes.Module())
	return mb.Create()
}

// ModuleDevtools ...
func ModuleForDevtools() application.Module {
	mb := libgin.NewDevtoolsModule()
	mb.Components(devtools4libgin.ExportComponents)

	mb.Depend(Module())

	return mb.Create()
}

// ModuleDemo ...
func ModuleForDemo() application.Module {
	mb := libgin.NewDemoModule()
	mb.Components(demo4libgin.ExportComponents)

	mb.Depend(ModuleForDevtools())
	mb.Depend(mimetypes.ModuleForCommon())

	return mb.Create()
}

// ModuleDemo ...
func ModuleForTest() application.Module {
	mb := libgin.NewTestModule()
	mb.Components(test4libgin.ExportComponents)

	mb.Depend(ModuleForDevtools())
	mb.Depend(mimetypes.ModuleForCommon())

	return mb.Create()
}
