package gen4devtools
import (
    p0ef6f2938 "github.com/starter-go/application"
    pd1a916a20 "github.com/starter-go/libgin"
    p9e4b55c70 "github.com/starter-go/libgin/devtools/controllers"
    p0f3e25981 "github.com/starter-go/libgin/devtools/groups"
     "github.com/starter-go/application"
)

// type p9e4b55c70.ArgumentController in package:github.com/starter-go/libgin/devtools/controllers
//
// id:com-9e4b55c707e952c9-controllers-ArgumentController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p9e4b55c707_controllers_ArgumentController struct {
}

func (inst* p9e4b55c707_controllers_ArgumentController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9e4b55c707e952c9-controllers-ArgumentController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9e4b55c707_controllers_ArgumentController) new() any {
    return &p9e4b55c70.ArgumentController{}
}

func (inst* p9e4b55c707_controllers_ArgumentController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9e4b55c70.ArgumentController)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)
    com.Responder = inst.getResponder(ie)


    return nil
}


func (inst*p9e4b55c707_controllers_ArgumentController) getContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p9e4b55c707_controllers_ArgumentController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}



// type p9e4b55c70.ComponentController in package:github.com/starter-go/libgin/devtools/controllers
//
// id:com-9e4b55c707e952c9-controllers-ComponentController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p9e4b55c707_controllers_ComponentController struct {
}

func (inst* p9e4b55c707_controllers_ComponentController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9e4b55c707e952c9-controllers-ComponentController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9e4b55c707_controllers_ComponentController) new() any {
    return &p9e4b55c70.ComponentController{}
}

func (inst* p9e4b55c707_controllers_ComponentController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9e4b55c70.ComponentController)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)
    com.Responder = inst.getResponder(ie)


    return nil
}


func (inst*p9e4b55c707_controllers_ComponentController) getContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p9e4b55c707_controllers_ComponentController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}



// type p9e4b55c70.EnvironmentController in package:github.com/starter-go/libgin/devtools/controllers
//
// id:com-9e4b55c707e952c9-controllers-EnvironmentController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p9e4b55c707_controllers_EnvironmentController struct {
}

func (inst* p9e4b55c707_controllers_EnvironmentController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9e4b55c707e952c9-controllers-EnvironmentController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9e4b55c707_controllers_EnvironmentController) new() any {
    return &p9e4b55c70.EnvironmentController{}
}

func (inst* p9e4b55c707_controllers_EnvironmentController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9e4b55c70.EnvironmentController)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)
    com.Responder = inst.getResponder(ie)


    return nil
}


func (inst*p9e4b55c707_controllers_EnvironmentController) getContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p9e4b55c707_controllers_EnvironmentController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}



// type p9e4b55c70.ExampleController in package:github.com/starter-go/libgin/devtools/controllers
//
// id:com-9e4b55c707e952c9-controllers-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p9e4b55c707_controllers_ExampleController struct {
}

func (inst* p9e4b55c707_controllers_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9e4b55c707e952c9-controllers-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9e4b55c707_controllers_ExampleController) new() any {
    return &p9e4b55c70.ExampleController{}
}

func (inst* p9e4b55c707_controllers_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9e4b55c70.ExampleController)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)
    com.Responder = inst.getResponder(ie)


    return nil
}


func (inst*p9e4b55c707_controllers_ExampleController) getContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p9e4b55c707_controllers_ExampleController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}



// type p9e4b55c70.PropertyController in package:github.com/starter-go/libgin/devtools/controllers
//
// id:com-9e4b55c707e952c9-controllers-PropertyController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p9e4b55c707_controllers_PropertyController struct {
}

func (inst* p9e4b55c707_controllers_PropertyController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9e4b55c707e952c9-controllers-PropertyController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9e4b55c707_controllers_PropertyController) new() any {
    return &p9e4b55c70.PropertyController{}
}

func (inst* p9e4b55c707_controllers_PropertyController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9e4b55c70.PropertyController)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)
    com.Responder = inst.getResponder(ie)


    return nil
}


func (inst*p9e4b55c707_controllers_PropertyController) getContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p9e4b55c707_controllers_PropertyController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}



// type p0f3e25981.DevtoolsGroup in package:github.com/starter-go/libgin/devtools/groups
//
// id:com-0f3e25981e100286-groups-DevtoolsGroup
// class:class-d1a916a203352fd5d33eabc36896b42e-Group
// alias:
// scope:singleton
//
type p0f3e25981e_groups_DevtoolsGroup struct {
}

func (inst* p0f3e25981e_groups_DevtoolsGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0f3e25981e100286-groups-DevtoolsGroup"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Group"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0f3e25981e_groups_DevtoolsGroup) new() any {
    return &p0f3e25981.DevtoolsGroup{}
}

func (inst* p0f3e25981e_groups_DevtoolsGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0f3e25981.DevtoolsGroup)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)


    return nil
}


func (inst*p0f3e25981e_groups_DevtoolsGroup) getContext(ie application.InjectionExt)pd1a916a20.Context{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Context").(pd1a916a20.Context)
}


