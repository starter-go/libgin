package main

import (
	"os"

	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/starter"
)

func main() {
	a := os.Args
	m := libgin.ModuleForTest()
	i := starter.Init(a)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
