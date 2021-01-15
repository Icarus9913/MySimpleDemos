package main

import (
	"fmt"
	"reflect"
)

func main() {

	shit()
}

func myKind() {
	var a func(i, j int) (string, error)
	reflect.TypeOf(a)        		//a的类型--->func(int, int) (string, error)
	reflect.TypeOf(a).Kind() 		//a具体的类型--->func

	if reflect.TypeOf(a).Kind() != reflect.Func {
		fmt.Println("不是")
	} else {
		fmt.Println("是") 		//reflect.Func是func.因此输出是
	}
}

func myElem() {
	type person struct{}
	p := &person{}
	reflect.TypeOf(p)       		 	//*main.person
	reflect.TypeOf(p).Kind()		 	//ptr
	reflect.TypeOf(p).Elem() 			//main.person.
	// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.

	pp := person{}
	reflect.TypeOf(pp)        //main.person
	reflect.TypeOf(pp).Kind() //struct
	//reflect.TypeOf(pp).Elem() ===>panic!!!!!
}

func myFuncOf() {
	var ta reflect.Type = reflect.ArrayOf(5, reflect.TypeOf(123))             // [5]int
	var tc reflect.Type = reflect.ChanOf(reflect.SendDir, ta)                          // chan<- [5]int
	var tp reflect.Type = reflect.PtrTo(ta)                                            // *[5]int
	var tf reflect.Type = reflect.FuncOf([]reflect.Type{ta}, []reflect.Type{tp, tc}, false) 	// func([5]int) (*[5]int, chan<- [5]int)

	fmt.Println(tf)
}

/*
	func MakeFunc(typ Type, fn func(args []Value) (results []Value)) Value
	-参数列表:
		typ Type 一个未初化函数的方法值，类型是reflect.Type
		fn func(args []Value) (results []Value) 另一个函数，作用于对第一个函数参数操作。
	-返回值:
		Value 返回 reflect.Value 类型，更多方法可以查看reflect.Value 结构中绑定的方法
	-功能说明:
		MakeFunc 返回一个新的类型“函数”包含 fn 函数（绑定着fn函数
*/
func myMakeFunc() {
	var swap = func(in []reflect.Value) []reflect.Value {
		fmt.Println("交换前:",in[0].Interface(),in[1].Interface())	//1 2
		return []reflect.Value{in[1],in[0]}
	}

	var makeSwap = func(fptr interface{}) {
		var valueOf reflect.Value = reflect.Indirect(reflect.ValueOf(fptr))
		var v reflect.Value = reflect.MakeFunc(valueOf.Type(),swap)
		valueOf.Set(v)

		fmt.Println(&valueOf)	//<func(int, int) (int, int) Value>  其中valueof与v的地址相同
		fmt.Println(&v)			//<func(int, int) (int, int) Value>
	}

	var intSwap func(int,int) (int,int)
	makeSwap(&intSwap)
	fmt.Println(intSwap(1,2))	//2 1
}

func shit()  {
	type invoker func(i,j int) int
	a := (invoker)(nil)
	b := (invoker)(nil,nil)

	fmt.Println(reflect.TypeOf(a))			//main.invoker
	fmt.Println(reflect.TypeOf(a).Kind())	//func

/*	makeFunc := reflect.MakeFunc(reflect.TypeOf((invoker)(nil)),
		func(args []reflect.Value) (results []reflect.Value) {

			fmt.Println("打印一下", args[0])
		return []reflect.Value{args[1]}
	})*/



}

