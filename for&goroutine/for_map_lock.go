package main

import (
	"math/rand"
	"sync"
)

//FIx the issue below to avoid "concurrent map writes" error
const NNN = 10

func main()  {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	wg.Add(NNN)
	for i:=0 ;i<NNN;i++{
		go func() {
			defer wg.Done()
			m[rand.Int()]=rand.Int()
		}()
	}
	wg.Wait()
	println(len(m))
}

//my fucking change
func change1()  {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	wg.Add(NNN)
	var l sync.Mutex
	for i:=0 ;i<NNN;i++{
		go func() {
			defer wg.Done()
			l.Lock()
			defer l.Unlock()
			m[rand.Int()]=rand.Int()
		}()
	}
	wg.Wait()
	println(len(m))
}
