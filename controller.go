package main

import (
	"github.com/hejiangda/diy-framework/framework/gin"
	"time"
)

//
//func FooControllerHandler(ctx *gin.Context) error {
//	durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), time.Duration(1*time.Second))
//	// 这里记得当所有事情处理结束后调用 cancel，告知 durationCtx 的后续 Context 结束
//	defer cancel()
//	finish := make(chan struct{}, 1)
//
//	panicChan := make(chan interface{}, 1)
//
//	go func() {
//		defer func() {
//			if p := recover(); p != nil {
//				panicChan <- p
//			}
//		}()
//
//		// 这里做具体的业务
//		time.Sleep(10 * time.Second)
//		ctx.ISetOkStatus().IJson("ok")
//
//		// 新的 goroutine 结束的时候通过一个 finish 通道告知父 goroutine
//		finish <- struct{}{}
//	}()
//
//	select {
//	case p := <-panicChan:
//		ctx.WriterMux().Lock()
//		defer ctx.WriterMux().Unlock()
//		ctx.SetStatus(http.StatusInternalServerError).Json("panic")
//		fmt.Println(p)
//	case <-finish:
//		fmt.Println("finish")
//	case <-durationCtx.Done():
//		ctx.WriterMux().Lock()
//		defer ctx.WriterMux().Unlock()
//		ctx.SetStatus(http.StatusInternalServerError).Json("time out")
//		ctx.SetHasTimeout()
//	}
//	return nil
//}

func UserLoginController(ctx *gin.Context) {
	time.Sleep(10 * time.Second)
	ctx.ISetOkStatus().IJson("ok, UserLoginController")
}
func SubjectDelController(ctx *gin.Context) {
	ctx.ISetOkStatus().IJson("ok, SubjectDelController")
}
func SubjectUpdateController(ctx *gin.Context) {
	ctx.ISetOkStatus().IJson("ok, SubjectUpdateController")
}
func SubjectListController(ctx *gin.Context) {
	ctx.ISetOkStatus().IJson("ok, SubjectListController")
}
func SubjectGetController(ctx *gin.Context) {
	ctx.ISetOkStatus().IJson("ok, SubjectGetController")
}
