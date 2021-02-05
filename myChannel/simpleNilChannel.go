package main

import (
	"fmt"
	"math/rand"
)

/*
	死锁是因为主线程被堵住了,无法继续往下走了.
	若是在main的for循环里加上一个time.After跳出来.就表明整个程序有某一条路是可以走通的.
	而go mySend(abc)这里表示在后台中等待中处于阻塞状态.
	==值为nil的通道,是特殊类型,因为它们总是阻塞的==
*/

func main() {
	var abc chan int
	go mySend(abc)

	for {
		select {
		case <-abc:
			fmt.Println("done")
		}
	}
}

func myWrite(abc chan int)  {

}


func mySend(abc chan int) {
	abc <- rand.Intn(10)
}
