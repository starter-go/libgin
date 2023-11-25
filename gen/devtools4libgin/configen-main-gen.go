package devtools4libgin

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p0f3e25981e_groups_DevtoolsGroup{})
    inst.register(&p9e4b55c707_controllers_ArgumentController{})
    inst.register(&p9e4b55c707_controllers_ComponentController{})
    inst.register(&p9e4b55c707_controllers_EnvironmentController{})
    inst.register(&p9e4b55c707_controllers_ExampleController{})
    inst.register(&p9e4b55c707_controllers_PropertyController{})


    return nil
}
