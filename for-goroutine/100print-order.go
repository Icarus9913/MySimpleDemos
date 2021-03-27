package main

import (
	"fmt"
	"time"
)

/*
	启动100个goroutine，并使其顺序的打印出从0-99数字，以下有两种方式
*/

func main() {
	planB()
}

//方法1
func planA() {
	go func() {
		c[0] <- 1
	}()

	run(100)
	time.Sleep(3 * time.Second)
}

/*
	功能：定义一个chan int的切片c，并对其中的每个chan int进行初始化，dart写法，
		 注意先初始化chan的切片，再对每个chan初始化
*/
var c = func() []chan int {
	v := make([]chan int, 100)	//这里的100定义的是切片的length
	for k := range v {
		v[k] = make(chan int, 0)	//这里的0是chan int无缓冲
	}
	return v
}()

/*
	启动100个goroutine，每个goroutine都等待着阻塞，意思是，
	每个goroutine中对于c这个chan切片，从0-99个chan int，顺序接收，
	这里打印的是for循环的i值.
	注意，此方案是的goroutine执行是顺序执行的，也就是会从100个goroutine里去找指定顺序的goroutine执行
*/
func run(n int) {
	for i := 0; i < n; i++ {
		temp := i
		go func(n int) {
			<-c[n]
			fmt.Println(n)
			if n+1 == 100 {
				return
			}
			c[n+1] <- 1
		}(temp)
	}
}


//---------------------------------------------------------------------

var d = make(chan int,0)	//无缓冲

func planB()  {
	go func() {
		d <- 0
	}()
	start(100)

	time.Sleep(3*time.Second)
}

/*
	启动100个goroutine，在主线程中给定义的chan全局变量d扔进去初始值0，
	然后每个goroutine接收这个chan变量，接着打印这个变量，然后再扔一个+1的值给这个全局chan变量
	注意，此时打印的goroutine无所谓是哪个，谁先接受到d就直接打印，反正后面再传入一个+1的值给d
*/
func start(n int)  {
	for i:=0;i<n;i++{
		go func() {
			temp := <-d
			fmt.Println(temp)
			d <-temp+1
		}()
	}
}
