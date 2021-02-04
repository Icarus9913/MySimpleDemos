package main

import (
	"fmt"
	"time"
)
var reada = make(chan int)

func readme() int {
	return <-reada
}

func main()  {
	go monitor1()
	time.Sleep(time.Millisecond)
	fmt.Println(readme())
}

func monitor1()  {
	var a int=100
	var index int
F:	for {
		select {
		case reada<-a:
			break F
		default:
			index++
		}
	}
	fmt.Println(index)
}


