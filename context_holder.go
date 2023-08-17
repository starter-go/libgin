package libgin

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

const theContextHolderKey = "libgin.ContextHolder#binding"

// ContextHolder 用来绑定 gin.Context & context.Context
type ContextHolder struct {
	gc *gin.Context
}

// GinContext 获取绑定的 gin.Context
func (inst *ContextHolder) GinContext() *gin.Context {
	return inst.gc
}

// GetContextHolder 取 ContextHolder
func GetContextHolder(cc context.Context) (*ContextHolder, error) {
	const key = theContextHolderKey
	o := cc.Value(key)
	h, ok := o.(*ContextHolder)
	if !ok {
		return nil, fmt.Errorf("no libgin.ContextHolder binding")
	}
	return h, nil
}

// BindContext 绑定 gin.Context & context.Context
func BindContext(c *gin.Context) *ContextHolder {
	const key = theContextHolderKey
	o := c.Value(key)
	h, ok := o.(*ContextHolder)
	if ok {
		return h
	}
	h = &ContextHolder{gc: c}
	c.Set(key, h)
	return h
}
