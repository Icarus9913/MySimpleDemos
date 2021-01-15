package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var a func(i,j int) (string,error)

	value := reflect.ValueOf(a)		//打印a的具体值
	fmt.Println(value)

	typeof := reflect.TypeOf(a)		//a的类型--->func(int, int) (string, error)
	fmt.Println(typeof)

	kind := typeof.Kind()			//a具体的类型--->func
	fmt.Println(kind)

	if kind!=reflect.Func{
		fmt.Println("不是")
	}else {
		fmt.Println("是")		//reflect.Func是func.因此输出是
	}

}

