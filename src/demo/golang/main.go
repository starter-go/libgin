package main

import (
	"os"

	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/starter"
)

func main() {
	mod := libgin.ModuleDemo()
	i := starter.Init(os.Args)
	i.MainModule(mod)
	i.WithPanic(true).Run()
}
