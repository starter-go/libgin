package libgin

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/base/context2"
)

const theContextHolderKey = "libgin.ContextHolder#binding"

////////////////////////////////////////////////////////////////////////////////

// ContextHolder 用来绑定 gin.Context & context.Context
type ContextHolder struct {
	gc *gin.Context
}

// GinContext 获取绑定的 gin.Context
func (inst *ContextHolder) GinContext() *gin.Context {
	return inst.gc
}

func (inst *ContextHolder) init() context2.Values {
	const key = theContextHolderKey
	context2.Setup(inst)
	inst.SetValue(key, inst)
	return inst
}

// GetValue ...
func (inst *ContextHolder) GetValue(key any) any {
	return inst.gc.Value(key)
}

// SetValue ...
func (inst *ContextHolder) SetValue(key, value any) {
	skey := key.(string)
	inst.gc.Set(skey, value)
}

// Context ...
func (inst *ContextHolder) Context() context.Context {
	return inst.gc
}

////////////////////////////////////////////////////////////////////////////////

// GetContextHolder 取 ContextHolder
func GetContextHolder(cc context.Context) (*ContextHolder, error) {
	const key = theContextHolderKey
	values, err := context2.GetValues(cc)
	if err != nil {
		return nil, err
	}
	h, ok := values.GetValue(key).(*ContextHolder)
	if !ok {
		return nil, fmt.Errorf("need invoke libgin.BindContext() first")
	}
	return h, nil
}

// BindContext 绑定 gin.Context & context.Context
func BindContext(c *gin.Context) *ContextHolder {
	const key = theContextHolderKey
	values, err := context2.GetValues(c)
	if values == nil || err != nil {
		h2 := &ContextHolder{gc: c}
		h2.init()
		values = h2
	}
	holder, ok := values.GetValue(key).(*ContextHolder)
	if !ok {
		panic("bad binding of libgin.ContextHolder")
	}
	return holder
}

////////////////////////////////////////////////////////////////////////////////
