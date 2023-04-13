package main

import "github.com/hejiangda/diy-framework/framework/gin"

func registerRouter(core *gin.Engine) {
	// 设置控制器

	//core.Get("foo", FooControllerHandler)
	core.GET("/user/login", UserLoginController)
	subjectApi := core.Group("/subject")
	{
		subjectApi.DELETE("/:id", SubjectDelController)
		subjectApi.PUT("/:id", SubjectUpdateController)
		subjectApi.GET("/:id", SubjectGetController)
		subjectApi.GET("/list/all", SubjectListController)
	}
}
