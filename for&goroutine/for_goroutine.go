package main

import "sync"

/*
	What will be printed when the code below is executed?
	And fix the issue to assure that len(m) is printed as 10
*/

const NN = 10

func main()  {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(NN)
	for i:=0;i<NN;i++{
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i]=i
			mu.Unlock()
		}()
	}
	wg.Wait()
	println(len(m))
}

//my fucking change
func change()  {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(NN)
	for i:=0;i<NN;i++{
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i]=i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	println(len(m))
}
