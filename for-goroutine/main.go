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
/*		p:=&v
		fmt.Println(unsafe.Pointer(&v))
		go p.print()*/

		go v.print()
	}

	time.Sleep(1 * time.Second)
}

/*
一. 需要注意的是,for range循环中,v这个变量的地址是一样的.只是值有改变而已,地址是复用的
二. 需要注意的另外一个问题是,func (f *field)print()函数接受*field类型对象的调用,还接受不同field类型对象的调用
三.
	1.当str是普通值的时候,对应v的类型是main.field, 调用go v.print的时候,调用的是v这个临时变量的地址去调用方法, v这个局部变量的地址是不变的,所以v实际是确定的,变的只是它存的值,
	  随着遍历结束,存的值可能已经变成最后的g了(还可能是其他的). 每次调用的是同一个对象同一个方法,只是for循环改变这个对象的值
	其中协程调用的时间消耗比for循环长,所以str普通值去调用print方法时候,会出现很多ggggggggggg输出,其中也会夹杂着个别其他的例如a,c其他的输出

	2.当str是指针值的时候,对应v的类型是*main.field,调用go v.print的时候,调用的是v的值(str数组的指针)所对应的方法,每次都是不同的对象的同一个方法,for循环改变对象.
	  就算for循环变成了&{g},但它之前调用的不是这个.

	3.如果print方法是func (f field)print的时候,这里str是普通值或者指针都可以.  当v是指针的时候,会去解引用调用方法.因为v是指针,他不能直接去调用,而是以v的指针所对应的数据去调用

	4.指针部分可以理解为go调用那一刻就已经确定了它要打印的内容,你遍历改变v和它没关系了.非指针部分调用那一刻不能确定要打印的内容,因为遍历会改变这个打印的内容
*/

