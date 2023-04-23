package http

import (
	"github.com/hejiangda/diy-framework/framework"
	"github.com/hejiangda/diy-framework/framework/gin"
)

// NewHttpEngine is command
func NewHttpEngine(container framework.Container) (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.SetContainer(container)
	r.Use(gin.Recovery())

	Routes(r)
	return r, nil
}
