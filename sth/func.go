package main

import "fmt"

type Human struct {
	Age int
}

type ss func(int) Human

func BigHello() func(int) ss {
	return func(int) ss {
		return func(i int) Human {
			fmt.Println("hello")
			var h Human
			h.Age = i
			return h
		}
	}
}

func main() {
	hello := BigHello()
	s := hello(1)
	s2 := s(1)
	fmt.Println(s2)

}
