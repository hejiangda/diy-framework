package main

import (
	"github.com/hejiangda/diy-framework/app/console"
	"github.com/hejiangda/diy-framework/app/http"
	"github.com/hejiangda/diy-framework/framework"
	"github.com/hejiangda/diy-framework/framework/provider/app"
	"github.com/hejiangda/diy-framework/framework/provider/config"
	"github.com/hejiangda/diy-framework/framework/provider/distributed"
	"github.com/hejiangda/diy-framework/framework/provider/env"
	"github.com/hejiangda/diy-framework/framework/provider/id"
	"github.com/hejiangda/diy-framework/framework/provider/kernel"
	"github.com/hejiangda/diy-framework/framework/provider/log"
	"github.com/hejiangda/diy-framework/framework/provider/trace"
)

func main() {
	// 初始化服务容器
	container := framework.NewHadeContainer()
	// 绑定App服务提供者
	container.Bind(&app.HadeAppProvider{})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.HadeEnvProvider{})
	container.Bind(&distributed.LocalDistributedProvider{})
	container.Bind(&config.HadeConfigProvider{})
	container.Bind(&id.HadeIDProvider{})
	container.Bind(&trace.HadeTraceProvider{})
	container.Bind(&log.HadeLogServiceProvider{})

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}
	// 运行root命令
	console.RunCommand(container)
}

//func main() {
//core := gin.New()
//core.Bind(&demo.DemoServiceProvider{})
//core.Bind(&app.HadeAppProvider{})
//
//core.Use(middleware.Cost())
//core.Use(gin.Recovery())
//hadeHttp.Routes(core)
////registerRouter(core)
//server := &http.Server{
//	Handler: core,
//	Addr:    ":8888",
//}
//// 这个goroutine是启动服务的goroutine
//go func() {
//	server.ListenAndServe()
//}()
//
//// 当前的goroutine等待信号量
//quit := make(chan os.Signal)
//// 监控信号：SIGINT, SIGTERM, SIGQUIT
//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
//// 这里会阻塞当前goroutine等待信号
//<-quit
//
//// 调用Server.Shutdown graceful结束
//timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//defer cancel()
//
//if err := server.Shutdown(timeoutCtx); err != nil {
//	log.Fatal("Server Shutdown:", err)
//}
//}
