package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	该例子运行流程:
		send()函数一直给d赋值,  add()函数在1秒时间内一直读d的值,
		当超时后,走<-t.C,给这个d赋nil,此后这个nil的channel d一直阻塞. 所以send()无法走, input:=<-c也无法走
		而t这个timer里的C已经被<-t.C一次. 所以在add()里的for循环的select里没有一个case走下去
		最后main协程结束
*/

func add(d chan int) {
	sum := 0
	t := time.NewTimer(time.Second)
	for {
		select {
		case input := <-d:
			sum = sum + input
		case <-t.C:
			d = nil
			fmt.Println(sum)
		}
	}
}

func send(d chan int) {
	for {
		d <- rand.Intn(10)
	}
}

func main() {
	d := make(chan int)
	go add(d)
	go send(d)
	time.Sleep(3 * time.Second)
}
