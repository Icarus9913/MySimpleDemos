package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	TestD()

}



//------------------TestA---------------------
func generator(n int) <-chan int {
	inCh := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			inCh <- i
		}
		defer close(inCh)
	}()
	return inCh
}

func do(inCh <-chan int, outCh chan<- int) {
	for v := range inCh {
		fmt.Println("  v:", v)
		outCh <- v
	}
	//defer close(outCh)
}

func TestA() {
	inCh := generator(100)
	outCh := make(chan int, 10)
	for i := 0; i < 5; i++ {
		go do(inCh, outCh)
	}
	go do(inCh, outCh)

	var index int = 1

	for r := range outCh {
		fmt.Println("---index:", index)
		fmt.Println(r)
		index++
	}
	fmt.Println("~~~~final:", index)
}


//-------------------------TestB---------------------------
func TestB() {
	gohou := Gohou()
	for {
		select {
		case evt, ok := <-gohou:
			if ok {
				fmt.Println(evt)
			} else {
				fmt.Println("结束了")
				return
			}
		}
	}
}

func Gohou() chan int {
	events := make(chan int)
	go LoopAdd(events)
	return events
}

func LoopAdd(events chan int) {
	defer close(events)
	for i := 1; i <= 10; i++ {
		events <- i
	}
}


//--------------------------TestC-----------------------------
func TestC() {
	var wg sync.WaitGroup
	ss := Server{make(chan struct{},1),wg}
	go ss.Stop()
	go serverHandler(&ss)
	time.Sleep(time.Second)
}

type Server struct {
	serverStopChan chan struct{}
	stopWg         sync.WaitGroup
}

func (s *Server) Stop() {
	if s.serverStopChan == nil {
		panic("gorpc.Server: server must be started before stopping it")
	}
	//close(s.serverStopChan)
	//s.stopWg.Wait()
	//s.serverStopChan = nil
	s.serverStopChan <-struct{}{}
}

func serverHandler(s *Server)  {
	var index int
	for {
		//if i==5{
		//	runtime.Gosched()
		//}
		select {
		case <-s.serverStopChan:
			fmt.Println("done")
			fmt.Println(index)
			return
		default:
			//fmt.Println("nihao")
			index++
		}
	}
}

func unBlockRead(ch chan int) (x int, err error) {
	select {
	case x = <-ch:
		return x, nil
	case <-time.After(time.Microsecond):
		return 0, errors.New("read time out")
	}
}

func unBlockWrite(ch chan int, x int) (err error) {
	select {
	case ch <- x:
		return nil
	case <-time.After(time.Microsecond):
		return errors.New("read time out")
	}
}


//------------------------TestD------------------------
func TestD()  {
	intChan := make(chan int, 5)
	exitChan := make(chan bool, 1)
	exitChan2 := make(chan bool, 1)

	go read(intChan, exitChan, exitChan2)
	go write(intChan, exitChan, exitChan2)
	time.Sleep(2*time.Second)
}

func write(intChan chan int, exitChan, exitChan2 chan bool) {
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	close(intChan)
	<-exitChan2
	fmt.Println("aaaaaaaaaaaaaaaaaa")
	exitChan <- true
}

func read(intChan chan int, exitChan, exitChan2 chan bool) {
	for {
		i, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println(i)
	}
	exitChan2 <- true
	<-exitChan
	fmt.Println("bbbbbbbbbbbbbbb")
}