package libgin

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/starter"
)

const (
	theModuleName     = "github.com/starter-go/libgin"
	theModuleVersion  = "v1.0.11"
	theModuleRevision = 12

	theMainModuleResPath     = "src/main/resources"
	theDemoModuleResPath     = "src/demo/resources"
	theDevtoolsModuleResPath = "src/devtools/resources"
)

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/demo/resources"
var theDemoModuleResFS embed.FS

//go:embed "src/devtools/resources"
var theDevtoolsModuleResFS embed.FS

// NewMainModule 导出模块 [github.com/starter-go/libgin]
func NewMainModule() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)

	mb.Depend(starter.Module())
	return mb
}

// NewDevtoolsModule 导出模块 [github.com/starter-go/libgin#devtools]
func NewDevtoolsModule() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#devtools")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theDevtoolsModuleResFS, theDevtoolsModuleResPath)

	mb.Depend(starter.Module())
	return mb
}

// NewDemoModule 创建模块 [github.com/starter-go/libgin#demo]
func NewDemoModule() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#demo")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theDemoModuleResFS, theDemoModuleResPath)

	mb.Depend(starter.Module())
	return mb
}
