package main

import (
	"os"

	"github.com/starter-go/libgin/modulegin"
	"github.com/starter-go/starter"
)

func main() {
	i := starter.Init(os.Args)
	i.MainModule(modulegin.ModuleDevtools())
	i.WithPanic(true).Run()
}
