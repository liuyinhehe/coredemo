package main

import (
	"context"
	"coredemo/framework"
	"fmt"
	"log"
	"time"
)

//为单个请求设置超时 分三步
func FooControllerHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)       //负责通知结束
	panicChan := make(chan interface{}, 1) //负责通知panic异常
	//1.继承 request 的 Context，创建出一个设置超时时间的 Context；这里设置超时时间为1s
	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
	// 这里记得当所有事情处理结束后调用 cancel，告知 durationCtx 的后续 Context 结束
	defer cancel()

	// mu := sync.Mutex{}
	//第二步创建一个新的 Goroutine 来处理业务逻辑： finish := make(chan struct{}, 1)
	/*
		其实在 Golang 的设计中，每个 Goroutine 都是独立存在的，父 Goroutine 一旦使用 Go 关键字开启了一个
		子 Goroutine，父子 Goroutine 就是平等存在的，他们互相不能干扰。而在异常面前，所有 Goroutine 的异
		常都需要自己管理，不会存在父 Goroutine 捕获子 Goroutine 异常的操作。所以切记：在 Golang 中，每个
		 Goroutine 创建的时候，我们要使用 defer 和 recover 关键字为当前 Goroutine 捕获 panic 异常，并进
		 行处理，否则，任意一处 panic 就会导致整个进程崩溃！
	*/
	go func() {
		//增加异常处理
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// Do real action
		time.Sleep(10 * time.Second)
		c.Json(200, "ok")

		//新的 goroutine 结束的时候通过一个 finish 通道告知父 goroutine
		finish <- struct{}{}
	}()

	/*
		3.设计事件处理顺序，当前 Goroutine 监听超时时间 Contex 的 Done() 事件，和具体的业务处理结束
		事件，哪个先到就先处理哪个。
	*/
	select {
	//请求监听的时候增加锁机制，防止这个时候其他go routine也要操作responseWriter
	case p := <-panicChan: /*异常事件*/
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.Json(500, "panic")
	case <-finish: /*结束事件*/
		fmt.Println("finish")
	case <-durationCtx.Done(): /*超时事件*/
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "time out") //告知浏览器前端，返回一个字符串信息
		c.SetHasTimeout()       //超时标记位为ture，已经有输出，不需要再进行任何的response设置
	}
	return nil
}
