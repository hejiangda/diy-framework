package http

import (
	"github.com/hejiangda/diy-framework/app/http/middleware/cors"
	"github.com/hejiangda/diy-framework/app/http/module/demo"
	"github.com/hejiangda/diy-framework/framework/contract"
	"github.com/hejiangda/diy-framework/framework/gin"
	ginSwagger "github.com/hejiangda/diy-framework/framework/middleware/gin-swagger"
	"github.com/hejiangda/diy-framework/framework/middleware/static"
	staticFiles "github.com/swaggo/files"
)

func Routes(r *gin.Engine) {

	//r.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))
	//r.Static("/dist/", "./dist/")

	container := r.GetContainer()
	configService := container.MustMake(contract.ConfigKey).(contract.Config)

	r.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))
	r.Use(cors.Default())

	if configService.GetBool("app.swagger") == true {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(staticFiles.Handler))
	}

	demo.Register(r)
}
