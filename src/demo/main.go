package main

import (
	"embed"
	"os"

	"github.com/starter-go/application"
	"github.com/starter-go/libgin/gen/gen4demo"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/starter"
)

const (
	theModuleName     = "github.com/starter-go/libgin#demo"
	theModuleVersion  = "1.0.0"
	theModuleRevision = 1
	theModuleResPath  = "resources"
)

//go:embed "resources"
var theModuleResFS embed.FS

func main() {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName).Version(theModuleVersion).Revision(theModuleRevision)
	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Components(gen4demo.ExportForDemo)
	mb.Depend(libgin.ModuleDevtools())

	mod := mb.Create()
	i := starter.Init(os.Args)
	i.MainModule(mod)
	i.WithPanic(true).Run()
}
