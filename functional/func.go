package main

import "fmt"

func main()  {
	var p Person
	p.name=Hello
	name := p.name("a", "b")
	fmt.Println(name)
}

type Person struct {
	name func(a,b string) string
}

func Hello(a,b string) string {
	str := a + b
	return str
}