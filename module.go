package libgin

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/starter"
)

const (
	theModuleName      = "github.com/starter-go/libgin"
	theModuleVersion   = "v1.0.2"
	theModuleRevision  = 3
	theModuleResPath   = "src/main/resources"
	theModuleResPathDT = "src/devtools/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

//go:embed "src/devtools/resources"
var theModuleResFSDT embed.FS

// NewModule 导出模块 [github.com/starter-go/libgin]
func NewModule() *application.ModuleBuilder {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName).Version(theModuleVersion).Revision(theModuleRevision)
	mb.EmbedResources(theModuleResFS, theModuleResPath)
	// mb.Components(gen4gin.ExportForGin)
	mb.Depend(starter.Module())
	return mb // .Create()
}

// NewModuleDevtools 导出模块 [github.com/starter-go/libgin#devtools]
func NewModuleDevtools() *application.ModuleBuilder {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#devtools").Version(theModuleVersion).Revision(theModuleRevision)
	mb.EmbedResources(theModuleResFSDT, theModuleResPathDT)
	// mb.Components(gen4devtools.ExportForDevtools)
	// mb.Depend(Module())
	mb.Depend(starter.Module())
	return mb // .Create()
}
