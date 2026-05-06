package libgin

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/base/context2"
)

const theContextAgentKey = "libgin.GinContextAgent#binding"

type GinContextAgent interface {
	Get(cc context.Context) (*gin.Context, error)
}

// ////////////////////////////////////////////////////////////////////////////////

// // ContextHolder 用来绑定 gin.Context & context.Context
// type ContextHolder struct {
// 	gc *gin.Context
// }

// // GinContext 获取绑定的 gin.Context
// func (inst *ContextHolder) GinContext() *gin.Context {
// 	return inst.gc
// }

// func (inst *ContextHolder) init() context2.Values {
// 	const key = theContextHolderKey
// 	context2.Setup(inst)
// 	inst.SetValue(key, inst)
// 	return inst
// }

// // GetValue ...
// func (inst *ContextHolder) GetValue(key any) any {
// 	return inst.gc.Value(key)
// }

// // SetValue ...
// func (inst *ContextHolder) SetValue(key, value any) {
// 	skey := key.(string)
// 	inst.gc.Set(skey, value)
// }

// // Context ...
// func (inst *ContextHolder) Context() context.Context {
// 	return inst.gc
// }

// ////////////////////////////////////////////////////////////////////////////////

// // GetContextHolder 取 ContextHolder
// func GetContextHolder(cc context.Context) (*ContextHolder, error) {
// 	const key = theContextHolderKey
// 	values, err := context2.GetValues(cc)
// 	if err != nil {
// 		return nil, err
// 	}
// 	h, ok := values.GetValue(key).(*ContextHolder)
// 	if !ok {
// 		return nil, fmt.Errorf("need invoke libgin.BindContext() first")
// 	}
// 	return h, nil
// }

// // BindContext 绑定 gin.Context & context.Context
// func BindContext(c *gin.Context) *ContextHolder {
// 	const key = theContextHolderKey
// 	values, err := context2.GetValues(c)
// 	if values == nil || err != nil {
// 		h2 := &ContextHolder{gc: c}
// 		h2.init()
// 		values = h2
// 	}
// 	holder, ok := values.GetValue(key).(*ContextHolder)
// 	if !ok {
// 		panic("bad binding of libgin.ContextHolder")
// 	}
// 	return holder
// }

////////////////////////////////////////////////////////////////////////////////

func GetGinContext(cc context.Context) (*gin.Context, error) {

	const key = theContextAgentKey
	value := cc.Value(key)
	agent, ok := value.(GinContextAgent)

	if !ok {
		return nil, fmt.Errorf("libgin.GetGinContext() : no GinContextAgent")
	}
	if agent == nil {
		return nil, fmt.Errorf("libgin.GetGinContext() : GinContextAgent is nil")
	}

	return agent.Get(cc)
}

func SetupGinContext(gc *gin.Context) (*gin.Context, error) {

	_, err := context2.Setup(gc, func(name string, value *context2.Context) {

		const agentKey = theContextAgentKey
		ada := new(innerGinContextAdapter)
		agent := new(innerGinContextAgentImpl)

		value.Adapter = ada

		gc.Set(name, value)
		gc.Set(agentKey, agent)
	})
	if err != nil {
		return nil, err
	}
	return gc, nil
}

////////////////////////////////////////////////////////////////////////////////

type innerGinContextAdapter struct {
}

// GetValue implements context2.Adapter.
func (inst *innerGinContextAdapter) GetValue(c *context2.Context, name any) any {
	cc := c.Raw
	return cc.Value(name)
}

// SetValue implements context2.Adapter.
func (inst *innerGinContextAdapter) SetValue(c *context2.Context, name any, value any) {

	cc := c.Raw
	gc, ok := cc.(*gin.Context)

	if !ok {
		panic("innerGinContextAdapter : bad gin.Context")
	}

	if gc == nil {
		panic("innerGinContextAdapter : gin.Context is nil")
	}

	nameStr := fmt.Sprint(name)
	gc.Set(nameStr, value)
	c.PutKey(name)
}

func (inst *innerGinContextAdapter) _impl() context2.Adapter {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerGinContextAgentImpl struct {
}

// Get implements GinContextAgent.
func (inst *innerGinContextAgentImpl) Get(cc context.Context) (*gin.Context, error) {

	ctx, err := context2.GetContext(cc)
	if err != nil {
		return nil, err
	}

	gc, ok := ctx.Raw.(*gin.Context)

	if !ok {
		return nil, fmt.Errorf("libgin.innerGinContextAgentImpl : no gin.Context")
	}
	if gc == nil {
		return nil, fmt.Errorf("libgin.innerGinContextAgentImpl : gin.Context is nil")
	}

	return gc, nil
}

func (inst *innerGinContextAgentImpl) _impl() GinContextAgent {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
