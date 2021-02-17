package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		defer close(temp)

		w.Wait()
	}()

	select {
	case <-temp:
		return false
	case <-time.After(t):
		return true
	}
}

//go run timeOut2.go 10000
func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need a time duration!")
		return
	}

	var w sync.WaitGroup
	w.Add(1)
	t, err := strconv.Atoi(arguments[1])
	if nil != err {
		fmt.Println(err)
		return
	}

	duration := time.Duration(int32(t)) * time.Millisecond
	fmt.Printf("Timeout period is %s\n", duration)

	if timeout(&w, duration) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}

	w.Done()
	if timeout(&w, duration) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}
}

func test() {

	//每隔一段时间触发一次
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for{
			<-ticker.C
			fmt.Println("ticker")
		}
	}()

	//只会触发一次
	timer := time.NewTimer(time.Second * 2)
	go func() {
		for{
			<-timer.C
			fmt.Println("timer")
		}
	}()
	time.Sleep(time.Second * 20)
}