package main4libgin

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

    
    inst.register(&p08eb425b0a_implements_ContentTypeResourceLoader{})
    inst.register(&p08eb425b0a_implements_DefaultContentTypeManager{})
    inst.register(&p08eb425b0a_implements_DefaultContext{})
    inst.register(&p08eb425b0a_implements_DefaultRouter{})
    inst.register(&p08eb425b0a_implements_HTTPConnector{})
    inst.register(&p08eb425b0a_implements_HTTPSConnector{})
    inst.register(&p08eb425b0a_implements_JSONGinResponder{})
    inst.register(&p08eb425b0a_implements_MainResponder{})
    inst.register(&p08eb425b0a_implements_RESTGroupRegistry{})
    inst.register(&p08eb425b0a_implements_StaticController{})
    inst.register(&p08eb425b0a_implements_StaticGroup{})


    return nil
}
