package main

import (
	"fmt"
	"reflect"
)

func main() {
	var shit = []interface{}{
		0: A,
		1: B,
		2: C,
	}

	for _, y := range shit {
		valueOf := reflect.ValueOf(y)
		t := valueOf.Type() //例如func(int, float64) int
		in := t.In(0)       //函数A的第0的参数的类型
		out := t.Out(0)     //函数A的第0个返回值类型

		fmt.Println("in是", in)   //输出in是int
		fmt.Println("out是", out) //输出out是int
	}

	shit[0] = B(1) //此时有入参,因此是调用函数
	fmt.Println("打印一下: ", shit[0])
}

func A(i int, j float64) int {
	return i + 10
}

func B(i int) int {
	return i + 10
}

func C(i int) int {
	return i + 10
}
