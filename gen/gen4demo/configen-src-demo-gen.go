package gen4demo
import (
    pd1a916a20 "github.com/starter-go/libgin"
    p8597db6cf "github.com/starter-go/libgin/src/demo/demo1"
     "github.com/starter-go/application"
)

// type p8597db6cf.ExampleController in package:github.com/starter-go/libgin/src/demo/demo1
//
// id:com-8597db6cfe6e2929-demo1-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p8597db6cfe_demo1_ExampleController struct {
}

func (inst* p8597db6cfe_demo1_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8597db6cfe6e2929-demo1-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8597db6cfe_demo1_ExampleController) new() any {
    return &p8597db6cf.ExampleController{}
}

func (inst* p8597db6cfe_demo1_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8597db6cf.ExampleController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)


    return nil
}


func (inst*p8597db6cfe_demo1_ExampleController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


