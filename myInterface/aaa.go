package main

import "fmt"

type Sayer interface {
	say()
}

type Animal struct {
	*Cat
	*Dog
}

type Cat struct {
	CC int
}
type Dog struct {
	DD int
}

func (a *Animal)say()  {
	switch {
	case a.Cat!=nil:
		fmt.Println("喵喵")
	case a.Dog!=nil:
		fmt.Println("汪汪")
	}
}

func main()  {
	var handle Sayer
	handle=&Animal{&Cat{},nil}
	handle.say()
}




