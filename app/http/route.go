package http

import (
	"github.com/hejiangda/diy-framework/app/http/module/demo"
	"github.com/hejiangda/diy-framework/framework/gin"
	"github.com/hejiangda/diy-framework/framework/middleware/static"
)

func Routes(r *gin.Engine) {

	r.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))
	//r.Static("/dist/", "./dist/")

	demo.Register(r)
}
