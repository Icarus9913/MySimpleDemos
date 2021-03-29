package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

//本文讲解如何实现Go超时控制,以及goroutine泄漏,还有panic的捕捉(一层层的recover,然后通过channel传回去)


/*
	前言:在 go 的函数中使用 defer + recover 不能捕获在当前函数里面新开的 goroutine 的 panic。
		在新的 goroutine 里面 panic 了即使上层有 recover 也会导致进程退出。如果只是调用其他的函数里panic了,还是可以捕捉.
		原因在于 panic 仅保证当前 goroutine 下的 defer 都会被调到，但不保证其他协程的defer也会调到。
		panic 发生时会先处理完当前goroutine已经defer挂上去的任务，执行完毕后再退出整个程序（注意是退出进程而不只是协程）。

	解决方案:因此在开新的 goroutine 的时候，一定要要注意这里，不要以为最外层有了 recover 程序就一定不会挂。
			新的 goroutine 里面需要自己处理 panic，或者在外层定义一个 channel ，
			goroutine 内部捕获到异常后把异常塞到这个 channel 中，上层监听到 goroutine 内部的异常后再在当前层进行 panic，
			这样就把底层的 panic 抛到了上层，交由上层的 recover 统一处理。
*/
func main()  {
	const total = 10
	var wg sync.WaitGroup
	wg.Add(total)
	now := time.Now()
	for i:=0;i<total;i++{
		go func() {
			defer func() {
				if p := recover();p!=nil{
					fmt.Println("oops, panic")
				}
			}()
			defer wg.Done()
			requestWork(context.Background(),"any")
		}()
	}
	wg.Wait()
	fmt.Println("elapsed:",time.Since(now))
	time.Sleep(20*time.Second)
	fmt.Println("number of goroutines:",runtime.NumGoroutine())
}

func hardWork(job interface{}) error {
	//time.Sleep(time.Second*10)
	panic("oops")
	return nil
}

func requestWork(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	/*
		此处done加缓冲,则可以避免goroutine的泄漏,当没有缓冲的时候,由于上下文的超时已经发生,则这个requestWork已经return了
		而hardWork那个协程却还在一直堵着,当主线程结束时,他还在堵着,此时的goroutine就泄漏了.
		所以加上缓冲时 done <- hardWork(job)不管在是否超时都能写入而不卡住goroutine.
	*/
	done := make(chan error,1)
	panicChan := make(chan interface{},1)
	go func() {
		defer func() {
			if p:=recover();nil!=p{
				panicChan <- p
			}
		}()
		done <- hardWork(job)
	}()

	select {
	case err := <-done:
		return err
	case p := <-panicChan:
		panic(p)
	case <-ctx.Done():
		return ctx.Err()
	}
}
