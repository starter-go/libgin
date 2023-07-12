package libgin

import "github.com/gin-gonic/gin"

// Router ...
type Router interface {
	Registration() *RouterRegistration
	Engine() *gin.Engine
}

// RouterRegistration ...
type RouterRegistration struct {
	Name   string
	Router Router
}
