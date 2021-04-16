package main

import (
	"errors"
	"fmt"
)

/*
	error这个小写的error是一个接口,下面的check函数里面return的结构体里用到errors.New函数返回了errorString{}这个结构体.
	所以是实现了这个error接口.
	可以搜一下"嵌入字段"这个词.
*/

type Error1 struct{ error }
type Error2 struct{ error }
type Error3 struct{ error }

func main() {
	err := check(2)
	switch err.(type) {
	case *Error1:
		fmt.Println("1")
	case *Error2:
		fmt.Println("2")
	case *Error3:
		fmt.Println("3")
	}
}

func check(i int) (err error) {
	if i == 1 {
		return &Error1{errors.New("bad1")}
	}
	if i == 2 {
		return &Error1{errors.New("bad2")}
	}
	if i == 3 {
		return &Error1{errors.New("bad3")}
	}
	return
}
