package main

import "diy-framework/framework"

func registerRouter(core *framework.Core) {
	// 设置控制器
	core.Get("foo", FooControllerHandler)
}
