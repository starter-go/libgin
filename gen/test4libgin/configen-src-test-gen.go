package test4libgin
import (
    p0a849518c "github.com/starter-go/libgin/src/test/golang/units4libgin"
     "github.com/starter-go/application"
)

// type p0a849518c.Unit1 in package:github.com/starter-go/libgin/src/test/golang/units4libgin
//
// id:com-0a849518c8547190-units4libgin-Unit1
// class:
// alias:
// scope:singleton
//
type p0a849518c8_units4libgin_Unit1 struct {
}

func (inst* p0a849518c8_units4libgin_Unit1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0a849518c8547190-units4libgin-Unit1"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0a849518c8_units4libgin_Unit1) new() any {
    return &p0a849518c.Unit1{}
}

func (inst* p0a849518c8_units4libgin_Unit1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0a849518c.Unit1)
	nop(ie, com)

	


    return nil
}


