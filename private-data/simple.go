package main

import (
	"fmt"
	"unsafe"
)

type MyTT struct {
	a int
}

func NewMyT() *MyTT {
	return &MyTT{
		a:22,
	}
}

func main()  {
	t := NewMyT()
	p := unsafe.Pointer(t)
	i := (*int)(p)
	fmt.Println(*i)
}