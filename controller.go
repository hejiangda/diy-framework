package main

import (
	"context"
	"diy-framework/framework"
	"fmt"
	"net/http"
	"time"
)

func FooControllerHandler(ctx *framework.Context) error {
	durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), time.Duration(1*time.Second))
	// 这里记得当所有事情处理结束后调用 cancel，告知 durationCtx 的后续 Context 结束
	defer cancel()
	finish := make(chan struct{}, 1)

	panicChan := make(chan interface{}, 1)

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		// 这里做具体的业务
		time.Sleep(10 * time.Second)
		ctx.SetOkStatus().Json("ok")

		// 新的 goroutine 结束的时候通过一个 finish 通道告知父 goroutine
		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		ctx.SetStatus(http.StatusInternalServerError).Json("panic")
		fmt.Println(p)
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		ctx.SetStatus(http.StatusInternalServerError).Json("time out")
		ctx.SetHasTimeout()
	}
	return nil
}

func UserLoginController(ctx *framework.Context) error {
	time.Sleep(10 * time.Second)
	ctx.SetOkStatus().Json("ok, UserLoginController")
	return nil
}
func SubjectDelController(ctx *framework.Context) error {
	ctx.SetOkStatus().Json("ok, SubjectDelController")
	return nil
}
func SubjectUpdateController(ctx *framework.Context) error {
	ctx.SetOkStatus().Json("ok, SubjectUpdateController")
	return nil
}
func SubjectListController(ctx *framework.Context) error {
	ctx.SetOkStatus().Json("ok, SubjectListController")
	return nil
}
func SubjectGetController(ctx *framework.Context) error {
	ctx.SetOkStatus().Json("ok, SubjectGetController")
	return nil
}
