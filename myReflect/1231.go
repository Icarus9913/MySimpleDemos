package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var shit = []interface{}{
		0:A,
		1:B,
		2:C,
	}

	for k,y := range shit{
		fmt.Println("K是",k)

		valueOf := reflect.ValueOf(y)
		t := valueOf.Type()
		in := t.In(0)
		out := t.Out(0)

		fmt.Println("in是",in)		//输出in是int
		fmt.Println("out是",out)	//输出out是int
	}

	shit[0] = B(1)
	fmt.Println(shit[0])


}

func A(i int) int {
	return i+10
}

func B(i int) int {
	return i+10
}

func C(i int) int {
	return i+10
}
