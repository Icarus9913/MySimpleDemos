package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	该例子运行流程:
		send()函数一直给c赋值,  add()函数在1秒时间内一直读c的值,
		当超时后,走<-t.C,给这个c赋nil,此后这个nil的channel c一直阻塞. 所以send()无法走, input:=<-c也无法走
		而t这个timer里的C已经被<-t.C一次. 所以在add()里的for循环的select里没有一个case走下去
		最后main协程结束
*/

func add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)
	for {
		select {
		case input := <-c:
			sum = sum + input
		case <-t.C:
			c = nil
			fmt.Println(sum)
		}
	}
}

func send(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}

func main() {
	c := make(chan int)
	go add(c)
	go send(c)
	time.Sleep(3 * time.Second)
}
