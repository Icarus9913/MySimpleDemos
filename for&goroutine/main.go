package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (f *field) print() {
	fmt.Println(f.name)
}

func main() {
	//str := []*field{{"ni"}, {"hao"}, {"a"},{"b"},{"c"},{"d"},{"e"},{"f"},{"g"}}
	str := []field{{"ni"},{"hao"},{"a"},{"b"},{"c"},{"d"},{"e"},{"f"},{"g"}}

	for _, v := range str {
		//fmt.Println(unsafe.Pointer(&v))

		//fmt.Println(v)
		go v.print()
	}

	time.Sleep(1 * time.Second)
}


