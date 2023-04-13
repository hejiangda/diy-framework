package http

import (
	"github.com/hejiangda/diy-framework/app/http/module/demo"
	"github.com/hejiangda/diy-framework/framework/gin"
)

func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
