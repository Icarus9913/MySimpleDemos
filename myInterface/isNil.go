package main

import (
	"fmt"
	"reflect"
)

func main() {
	demo3()
}

func demo1() {
	var v interface{}
	v = (*int)(nil)
	fmt.Println(v == nil)	//false
}

func demo2() {
	var data *byte
	var in interface{}

	fmt.Println(data, data == nil)	//<nil> true
	fmt.Println(in, in == nil)	//<nil> true

	in = data
	fmt.Println(in, in == nil)	//<nil> false
}

/*
	interface不是一个指针类型
	interface :
		1- runtime.iface 结构体:表示不包含任何方法的空接口,也称为empty interface
		2- runtime.eface 结构体:表示包含方法的接口

	type eface struct{
		_type	*_type
		data	unsafe.Pointer
	}

	type iface struct{
		tab		*itab
		data	unsafe.Pointer
	}

	会发现interface不是单纯的值,而是分为类型和值,所以传统认知的此nil并非彼nil,必须得类型和值都为nil的情况下,
	interface的nil判断才会为true
*/

//---------------------------------------------------------------------------------------------------

/*
	解决方案:
	利用反射来做nil的值判断,在反射中会有针对interface类型的特殊处理,最终输出结果是:true,达到效果

	其他方法:
		- 对值进行nil判断,再返回给interface设置
		- 返回具体的值的类型,而不是返回interface
*/
func demo3()  {
	var data *byte
	var in interface{}

	in = data
	fmt.Println(IsNil(in))
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr{
		return vi.IsNil()
	}
	return false
}
