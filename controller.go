package main

import (
	"context"
	"diy-framework/framework"
	"fmt"
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
		ctx.Json(200, "ok")

		// 新的 goroutine 结束的时候通过一个 finish 通道告知父 goroutine
		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		ctx.WriteMux().Lock()
		defer ctx.WriteMux().Unlock()
		ctx.Json(500, "panic")
		fmt.Println(p)
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		ctx.WriteMux().Lock()
		defer ctx.WriteMux().Unlock()
		ctx.Json(500, "time out")
		ctx.SetTimeout()
	}
	return nil
}
