package main

import "fmt"

func main() {
	c1 := receiver()
	for {
		select {
		case i:=<-c1:
			fmt.Println(i)
		}
	}

}

func receiver() <-chan int {
	c := make(chan int, 5)
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
	return c
}
