package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

//go run -race raceC.go 10
func main() {
	//var m sync.Mutex
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Give me a natural number!")
		os.Exit(1)
	}
	numberGR, err := strconv.Atoi(os.Args[1])
	if nil != err {
		fmt.Println(err)
		return
	}

	var waitGroup sync.WaitGroup
	var i int

	k := make(map[int]int)
	k[1] = 12

	for i = 0; i < numberGR; i++ {
		waitGroup.Add(1)
		go func(j int) {
			defer waitGroup.Done()
			m.Lock()
			k[j] = j
			m.Unlock()
		}(i)
	}
	waitGroup.Wait()

	k[2] = 10
	fmt.Printf("k = %v\n", k)
}
