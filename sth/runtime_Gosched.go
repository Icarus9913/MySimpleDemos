package main

import (
	"fmt"
	"runtime"
	"sync"
)

//Add code in line A to assure that the lowercase letters and capital letters are printed consecutively
const N = 26

func main()  {
	const GOMAXPROCS = 1
	runtime.GOMAXPROCS(GOMAXPROCS)

	var wg sync.WaitGroup
	wg.Add(2*N)
	for i:=0;i<N;i++{
		go func(i int) {
			defer wg.Done()
			//A
			fmt.Printf("%c",'a'+i)
		}(i)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%c",'A'+i)
		}(i)
	}
	wg.Wait()
}

//my fucking change
func WatchOut()  {
	const GOMAXPROCS = 1
	runtime.GOMAXPROCS(GOMAXPROCS)

	var wg sync.WaitGroup
	wg.Add(2*N)
	for i:=0;i<N;i++{
		go func(i int) {
			defer wg.Done()
			runtime.Gosched()
			fmt.Printf("%c",'a'+i)
		}(i)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%c",'A'+i)
		}(i)
	}
	wg.Wait()
}


