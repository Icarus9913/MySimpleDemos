package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan []byte,10)
	go func() {
		for {
			select {
			case d := <-ch:
				fmt.Println(string(d))
			}
		}
	}()

	data := make([]byte, 0, 32)
	data = append(data, []byte("bbbbbbbbbb")...)
	ch <- data

	data = data[:0]
	//time.Sleep(time.Second)
	data = append(data, []byte("aaa")...)
	ch <- data
	time.Sleep(time.Second * 2)
}

/*
 输出:
	aaabbbbbbb
	aaa
 解释:(丢进通道的是一个切片,有指针性质)
	第一次从channel拿到数据时,跟data指向同一底层地址.然后还运行到print时候,
	data = data[:0]把data清空,data又append(aaa),表示从头开始存入.
	接着print执行了,第一次打印结果就是aaabbbbbbbb
	第二次从channel拿到data的时候,已经变成aaa了

	ch := make(chan []byte,10)  或者把这里的缓冲给去掉,不要缓冲.此时就会读一个写一个,阻塞的读.
*/


