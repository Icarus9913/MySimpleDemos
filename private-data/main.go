package main

import (
	"fmt"
	"unsafe"
)

type TT struct {
	a int
	b int
}

func NewTT() *TT {
	return &TT{
		a:1,
		b:2,
	}
}

func main()  {
	t := NewTT()
	p := unsafe.Pointer(t)
	m := func() unsafe.Pointer {
		temp := uintptr(p)
		a := 1
		temp += unsafe.Sizeof(a)
		return unsafe.Pointer(temp)
	}()

	i := (*int)(m)
	fmt.Println(*i)
}
