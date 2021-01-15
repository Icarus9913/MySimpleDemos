package main

import (
	"fmt"
	"reflect"
)

func main() {

	var num float64 = 1.2345
	fmt.Println("old value of pointer:", num)		//old value of pointer: 1.2345

	// 通过reflect.ValueOf获取num中的reflect.Value，注意，参数必须是指针才能修改其值
	pointer := reflect.ValueOf(&num)
	newValue := pointer.Elem()							// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.

	fmt.Println("type of pointer:", newValue.Type())			//type of pointer: float64
	fmt.Println("settability of pointer:", newValue.CanSet())	//settability of pointer: true


	// 重新赋值
	newValue.SetFloat(77)
	fmt.Println("new value of pointer:", num)		//new value of pointer: 77

	////////////////////
	// 如果reflect.ValueOf的参数不是指针，会如何？
	pointer = reflect.ValueOf(num)
	//newValue = pointer.Elem() // 如果非指针，这里直接panic，“panic: reflect: call of reflect.Value.Elem on float64 Value”
}
