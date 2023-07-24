package gen4gin
import (
    p0ef6f2938 "github.com/starter-go/application"
    pd1a916a20 "github.com/starter-go/libgin"
    p08eb425b0 "github.com/starter-go/libgin/implements"
     "github.com/starter-go/application"
)

// type p08eb425b0.ContentTypeResourceLoader in package:github.com/starter-go/libgin/implements
//
// id:com-08eb425b0ad6e539-implements-ContentTypeResourceLoader
// class:class-d1a916a203352fd5d33eabc36896b42e-ContentTypeRegistry
// alias:
// scope:singleton
//
type p08eb425b0a_implements_ContentTypeResourceLoader struct {
}

func (inst* p08eb425b0a_implements_ContentTypeResourceLoader) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-08eb425b0ad6e539-implements-ContentTypeResourceLoader"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-ContentTypeRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p08eb425b0a_implements_ContentTypeResourceLoader) new() any {
    return &p08eb425b0.ContentTypeResourceLoader{}
}

func (inst* p08eb425b0a_implements_ContentTypeResourceLoader) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p08eb425b0.ContentTypeResourceLoader)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)
    com.TypesProps = inst.getTypesProps(ie)


    return nil
}


func (inst*p08eb425b0a_implements_ContentTypeResourceLoader) getContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p08eb425b0a_implements_ContentTypeResourceLoader) getTypesProps(ie application.InjectionExt)string{
    return ie.GetString("${web.content-types.properties}")
}



// type p08eb425b0.DefaultContentTypeManager in package:github.com/starter-go/libgin/implements
//
// id:com-08eb425b0ad6e539-implements-DefaultContentTypeManager
// class:
// alias:alias-d1a916a203352fd5d33eabc36896b42e-ContentTypeManager
// scope:singleton
//
type p08eb425b0a_implements_DefaultContentTypeManager struct {
}

func (inst* p08eb425b0a_implements_DefaultContentTypeManager) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-08eb425b0ad6e539-implements-DefaultContentTypeManager"
	r.Classes = ""
	r.Aliases = "alias-d1a916a203352fd5d33eabc36896b42e-ContentTypeManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p08eb425b0a_implements_DefaultContentTypeManager) new() any {
    return &p08eb425b0.DefaultContentTypeManager{}
}

func (inst* p08eb425b0a_implements_DefaultContentTypeManager) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p08eb425b0.DefaultContentTypeManager)
	nop(ie, com)

	
    com.All = inst.getAll(ie)


    return nil
}


func (inst*p08eb425b0a_implements_DefaultContentTypeManager) getAll(ie application.InjectionExt)[]pd1a916a20.ContentTypeRegistry{
    dst := make([]pd1a916a20.ContentTypeRegistry, 0)
    src := ie.ListComponents(".class-d1a916a203352fd5d33eabc36896b42e-ContentTypeRegistry")
    for _, item1 := range src {
        item2 := item1.(pd1a916a20.ContentTypeRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type p08eb425b0.DefaultContext in package:github.com/starter-go/libgin/implements
//
// id:com-08eb425b0ad6e539-implements-DefaultContext
// class:
// alias:alias-d1a916a203352fd5d33eabc36896b42e-Context
// scope:singleton
//
type p08eb425b0a_implements_DefaultContext struct {
}

func (inst* p08eb425b0a_implements_DefaultContext) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-08eb425b0ad6e539-implements-DefaultContext"
	r.Classes = ""
	r.Aliases = "alias-d1a916a203352fd5d33eabc36896b42e-Context"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p08eb425b0a_implements_DefaultContext) new() any {
    return &p08eb425b0.DefaultContext{}
}

func (inst* p08eb425b0a_implements_DefaultContext) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p08eb425b0.DefaultContext)
	nop(ie, com)

	
    com.Controllers = inst.getControllers(ie)
    com.Connectors = inst.getConnectors(ie)
    com.Routers = inst.getRouters(ie)
    com.Groups = inst.getGroups(ie)
    com.DefaultGroupName = inst.getDefaultGroupName(ie)


    return nil
}


func (inst*p08eb425b0a_implements_DefaultContext) getControllers(ie application.InjectionExt)[]pd1a916a20.Controller{
    dst := make([]pd1a916a20.Controller, 0)
    src := ie.ListComponents(".class-d1a916a203352fd5d33eabc36896b42e-Controller")
    for _, item1 := range src {
        item2 := item1.(pd1a916a20.Controller)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*p08eb425b0a_implements_DefaultContext) getConnectors(ie application.InjectionExt)[]pd1a916a20.Connector{
    dst := make([]pd1a916a20.Connector, 0)
    src := ie.ListComponents(".class-d1a916a203352fd5d33eabc36896b42e-Connector")
    for _, item1 := range src {
        item2 := item1.(pd1a916a20.Connector)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*p08eb425b0a_implements_DefaultContext) getRouters(ie application.InjectionExt)[]pd1a916a20.Router{
    dst := make([]pd1a916a20.Router, 0)
    src := ie.ListComponents(".class-d1a916a203352fd5d33eabc36896b42e-Router")
    for _, item1 := range src {
        item2 := item1.(pd1a916a20.Router)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*p08eb425b0a_implements_DefaultContext) getGroups(ie application.InjectionExt)[]pd1a916a20.Group{
    dst := make([]pd1a916a20.Group, 0)
    src := ie.ListComponents(".class-d1a916a203352fd5d33eabc36896b42e-Group")
    for _, item1 := range src {
        item2 := item1.(pd1a916a20.Group)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*p08eb425b0a_implements_DefaultContext) getDefaultGroupName(ie application.InjectionExt)string{
    return ie.GetString("${web.default-group-name}")
}



// type p08eb425b0.DefaultRouter in package:github.com/starter-go/libgin/implements
//
// id:com-08eb425b0ad6e539-implements-DefaultRouter
// class:class-d1a916a203352fd5d33eabc36896b42e-Router
// alias:
// scope:singleton
//
type p08eb425b0a_implements_DefaultRouter struct {
}

func (inst* p08eb425b0a_implements_DefaultRouter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-08eb425b0ad6e539-implements-DefaultRouter"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Router"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p08eb425b0a_implements_DefaultRouter) new() any {
    return &p08eb425b0.DefaultRouter{}
}

func (inst* p08eb425b0a_implements_DefaultRouter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p08eb425b0.DefaultRouter)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)
    com.Name = inst.getName(ie)
    com.GroupList = inst.getGroupList(ie)
    com.GinMode = inst.getGinMode(ie)


    return nil
}


func (inst*p08eb425b0a_implements_DefaultRouter) getContext(ie application.InjectionExt)pd1a916a20.Context{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Context").(pd1a916a20.Context)
}


func (inst*p08eb425b0a_implements_DefaultRouter) getName(ie application.InjectionExt)string{
    return ie.GetString("${web-router.default.name}")
}


func (inst*p08eb425b0a_implements_DefaultRouter) getGroupList(ie application.InjectionExt)string{
    return ie.GetString("${web-router.default.groups}")
}


func (inst*p08eb425b0a_implements_DefaultRouter) getGinMode(ie application.InjectionExt)string{
    return ie.GetString("${gin.mode}")
}



// type p08eb425b0.HTTPConnector in package:github.com/starter-go/libgin/implements
//
// id:com-08eb425b0ad6e539-implements-HTTPConnector
// class:class-d1a916a203352fd5d33eabc36896b42e-Connector
// alias:
// scope:singleton
//
type p08eb425b0a_implements_HTTPConnector struct {
}

func (inst* p08eb425b0a_implements_HTTPConnector) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-08eb425b0ad6e539-implements-HTTPConnector"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Connector"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p08eb425b0a_implements_HTTPConnector) new() any {
    return &p08eb425b0.HTTPConnector{}
}

func (inst* p08eb425b0a_implements_HTTPConnector) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p08eb425b0.HTTPConnector)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)
    com.Host = inst.getHost(ie)
    com.Port = inst.getPort(ie)
    com.Enabled = inst.getEnabled(ie)


    return nil
}


func (inst*p08eb425b0a_implements_HTTPConnector) getContext(ie application.InjectionExt)pd1a916a20.Context{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Context").(pd1a916a20.Context)
}


func (inst*p08eb425b0a_implements_HTTPConnector) getHost(ie application.InjectionExt)string{
    return ie.GetString("${server.http.host}")
}


func (inst*p08eb425b0a_implements_HTTPConnector) getPort(ie application.InjectionExt)int{
    return ie.GetInt("${server.http.port}")
}


func (inst*p08eb425b0a_implements_HTTPConnector) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${server.http.enabled}")
}



// type p08eb425b0.HTTPSConnector in package:github.com/starter-go/libgin/implements
//
// id:com-08eb425b0ad6e539-implements-HTTPSConnector
// class:class-d1a916a203352fd5d33eabc36896b42e-Connector
// alias:
// scope:singleton
//
type p08eb425b0a_implements_HTTPSConnector struct {
}

func (inst* p08eb425b0a_implements_HTTPSConnector) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-08eb425b0ad6e539-implements-HTTPSConnector"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Connector"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p08eb425b0a_implements_HTTPSConnector) new() any {
    return &p08eb425b0.HTTPSConnector{}
}

func (inst* p08eb425b0a_implements_HTTPSConnector) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p08eb425b0.HTTPSConnector)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)
    com.Enabled = inst.getEnabled(ie)
    com.Host = inst.getHost(ie)
    com.Port = inst.getPort(ie)
    com.KeyFile = inst.getKeyFile(ie)
    com.CertFile = inst.getCertFile(ie)


    return nil
}


func (inst*p08eb425b0a_implements_HTTPSConnector) getContext(ie application.InjectionExt)pd1a916a20.Context{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Context").(pd1a916a20.Context)
}


func (inst*p08eb425b0a_implements_HTTPSConnector) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${server.https.enabled}")
}


func (inst*p08eb425b0a_implements_HTTPSConnector) getHost(ie application.InjectionExt)string{
    return ie.GetString("${server.https.host}")
}


func (inst*p08eb425b0a_implements_HTTPSConnector) getPort(ie application.InjectionExt)int{
    return ie.GetInt("${server.https.port}")
}


func (inst*p08eb425b0a_implements_HTTPSConnector) getKeyFile(ie application.InjectionExt)string{
    return ie.GetString("${server.https.key-file}")
}


func (inst*p08eb425b0a_implements_HTTPSConnector) getCertFile(ie application.InjectionExt)string{
    return ie.GetString("${server.https.certificate-file}")
}



// type p08eb425b0.MainResponder in package:github.com/starter-go/libgin/implements
//
// id:com-08eb425b0ad6e539-implements-MainResponder
// class:
// alias:alias-d1a916a203352fd5d33eabc36896b42e-Responder
// scope:singleton
//
type p08eb425b0a_implements_MainResponder struct {
}

func (inst* p08eb425b0a_implements_MainResponder) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-08eb425b0ad6e539-implements-MainResponder"
	r.Classes = ""
	r.Aliases = "alias-d1a916a203352fd5d33eabc36896b42e-Responder"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p08eb425b0a_implements_MainResponder) new() any {
    return &p08eb425b0.MainResponder{}
}

func (inst* p08eb425b0a_implements_MainResponder) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p08eb425b0.MainResponder)
	nop(ie, com)

	
    com.Responders = inst.getResponders(ie)


    return nil
}


func (inst*p08eb425b0a_implements_MainResponder) getResponders(ie application.InjectionExt)[]pd1a916a20.Responder{
    dst := make([]pd1a916a20.Responder, 0)
    src := ie.ListComponents(".class-d1a916a203352fd5d33eabc36896b42e-Responder")
    for _, item1 := range src {
        item2 := item1.(pd1a916a20.Responder)
        dst = append(dst, item2)
    }
    return dst
}



// type p08eb425b0.RESTGroup in package:github.com/starter-go/libgin/implements
//
// id:com-08eb425b0ad6e539-implements-RESTGroup
// class:class-d1a916a203352fd5d33eabc36896b42e-Group
// alias:
// scope:singleton
//
type p08eb425b0a_implements_RESTGroup struct {
}

func (inst* p08eb425b0a_implements_RESTGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-08eb425b0ad6e539-implements-RESTGroup"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Group"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p08eb425b0a_implements_RESTGroup) new() any {
    return &p08eb425b0.RESTGroup{}
}

func (inst* p08eb425b0a_implements_RESTGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p08eb425b0.RESTGroup)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)
    com.Name = inst.getName(ie)
    com.Path = inst.getPath(ie)


    return nil
}


func (inst*p08eb425b0a_implements_RESTGroup) getContext(ie application.InjectionExt)pd1a916a20.Context{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Context").(pd1a916a20.Context)
}


func (inst*p08eb425b0a_implements_RESTGroup) getName(ie application.InjectionExt)string{
    return ie.GetString("${web-group.rest.name}")
}


func (inst*p08eb425b0a_implements_RESTGroup) getPath(ie application.InjectionExt)string{
    return ie.GetString("${web-group.rest.path}")
}



// type p08eb425b0.StaticController in package:github.com/starter-go/libgin/implements
//
// id:com-08eb425b0ad6e539-implements-StaticController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p08eb425b0a_implements_StaticController struct {
}

func (inst* p08eb425b0a_implements_StaticController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-08eb425b0ad6e539-implements-StaticController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p08eb425b0a_implements_StaticController) new() any {
    return &p08eb425b0.StaticController{}
}

func (inst* p08eb425b0a_implements_StaticController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p08eb425b0.StaticController)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)
    com.ResPath = inst.getResPath(ie)
    com.IndexNames = inst.getIndexNames(ie)
    com.Types = inst.getTypes(ie)


    return nil
}


func (inst*p08eb425b0a_implements_StaticController) getContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p08eb425b0a_implements_StaticController) getResPath(ie application.InjectionExt)string{
    return ie.GetString("${web-group.static.resources}")
}


func (inst*p08eb425b0a_implements_StaticController) getIndexNames(ie application.InjectionExt)string{
    return ie.GetString("${web-group.static.index-names}")
}


func (inst*p08eb425b0a_implements_StaticController) getTypes(ie application.InjectionExt)pd1a916a20.ContentTypeManager{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-ContentTypeManager").(pd1a916a20.ContentTypeManager)
}



// type p08eb425b0.StaticGroup in package:github.com/starter-go/libgin/implements
//
// id:com-08eb425b0ad6e539-implements-StaticGroup
// class:class-d1a916a203352fd5d33eabc36896b42e-Group
// alias:
// scope:singleton
//
type p08eb425b0a_implements_StaticGroup struct {
}

func (inst* p08eb425b0a_implements_StaticGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-08eb425b0ad6e539-implements-StaticGroup"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Group"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p08eb425b0a_implements_StaticGroup) new() any {
    return &p08eb425b0.StaticGroup{}
}

func (inst* p08eb425b0a_implements_StaticGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p08eb425b0.StaticGroup)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)
    com.Path = inst.getPath(ie)
    com.Name = inst.getName(ie)


    return nil
}


func (inst*p08eb425b0a_implements_StaticGroup) getContext(ie application.InjectionExt)pd1a916a20.Context{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Context").(pd1a916a20.Context)
}


func (inst*p08eb425b0a_implements_StaticGroup) getPath(ie application.InjectionExt)string{
    return ie.GetString("${web-group.static.path}")
}


func (inst*p08eb425b0a_implements_StaticGroup) getName(ie application.InjectionExt)string{
    return ie.GetString("${web-group.static.name}")
}


