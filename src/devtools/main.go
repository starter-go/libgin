package main

import (
	"os"

	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/starter"
)

func main() {
	i := starter.Init(os.Args)
	i.MainModule(libgin.ModuleForDevtools())
	i.WithPanic(true).Run()
}
