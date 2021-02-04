package main

import (
	"fmt"
	"sync"
)

var m1 sync.Mutex

func function()  {
	m1.Lock()
	fmt.Println("Locked!")
}

func main()  {
	var w sync.WaitGroup

	go func() {
		defer w.Done()
		function()
	}()
	w.Add(1)

	go func() {
		defer w.Done()
		function()
	}()
	w.Add(1)
	w.Wait()
}
