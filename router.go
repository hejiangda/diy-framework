package main

import (
	"diy-framework/framework"
)

func registerRouter(core *framework.Core) {
	// 设置控制器
	//core.Get("foo", FooControllerHandler)
	core.Get("/user/login", UserLoginController)
	subjectApi := core.Group("/subject")
	{
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)
	}
}
