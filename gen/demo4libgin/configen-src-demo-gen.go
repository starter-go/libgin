package demo4libgin
import (
    pd1a916a20 "github.com/starter-go/libgin"
    p4e2855b32 "github.com/starter-go/libgin/src/demo/golang/demo1"
     "github.com/starter-go/application"
)

// type p4e2855b32.ExampleController in package:github.com/starter-go/libgin/src/demo/golang/demo1
//
// id:com-4e2855b323169093-demo1-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p4e2855b323_demo1_ExampleController struct {
}

func (inst* p4e2855b323_demo1_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4e2855b323169093-demo1-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4e2855b323_demo1_ExampleController) new() any {
    return &p4e2855b32.ExampleController{}
}

func (inst* p4e2855b323_demo1_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4e2855b32.ExampleController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)


    return nil
}


func (inst*p4e2855b323_demo1_ExampleController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


