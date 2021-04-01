package main

import (
	"fmt"
	"time"
)

func main() {
	simple()
}

/*
	out是一个只声明但没有初始化的变量. 然后给他赋值一个已初始化的变量,则这两个chan变量都可以用
*/
func simple() {
	outCh := make(chan int)
	var out chan int
	out = outCh

	go func() {
		out <- 11
		outCh <- 22
	}()

	go func() {
		for v := range out { //或者range outCh也可以
			fmt.Println(v)
		}
	}()

	time.Sleep(time.Second)
}


/*
	先是53行收到值,执行完后该case就堵住了,因为in变成nil,
	然后瞬间发给50行. 再执行59行的func.
*/
func demo() {
	inCh := make(chan int)
	outCh := make(chan int)

	go func() {
		var in chan int = inCh
		var out chan int
		var val int

		for {
			select {
			case out <- val:
				out = nil
				in = inCh
			case val = <-in:
				out = outCh
				in = nil
			}
		}
	}()

	go func() {
		for r := range outCh{
			fmt.Println("Result: ",r)
		}
	}()

	inCh <- 1
	//inCh <- 2
	time.Sleep(2*time.Second)
}
