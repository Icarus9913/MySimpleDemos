package main

import (
	"fmt"
	"unsafe"
)

/*
	此结构体的对齐系数是d int64--->8
	所以,该结构体的总大小计算方式为: a为1个字节补3与b组成8,
	c下面的d是8,所以c单独需要组成8,e下面没了,所以e单独组成8
	总的就是:4+4+8+8+8=32
*/
type Part1 struct {
	a bool		//1byte
	b int32		//4
	c int8		//1
	d int64		//8
	e byte		//1
}

/*
	此结构体的对齐系数是d int64--->8
	所以,该结构体的总大小计算方式为: e与c与a组4,再与b结合组成8
	总的就是:4+4+8=16
*/
type Part2 struct {
	e byte		//1
	c int8		//1
	a bool		//1
	b int32		//4
	d int64		//8
}

/*
	此结构体的对齐系数是b int32--->4
	所以,该结构体的总大小计算方式为: e与c与a组4,再与b结合组成8
	总的就是:4+4=8
*/
type Part3 struct {
	e byte		//1
	c int8		//1
	a bool		//1
	b int32		//4
}

/*
	此结构体的对齐系数是c int64--->8
	其中空结构体size为0
	所以,该结构体的总大小计算方式为: d与a组4,再与b结合组成8
	总的就是:4+4+8=16
*/
type Part4 struct {
	d struct{}	//0
	a bool		//1
	b int32		//4
	c int64		//8
}

/*
	此结构体的对齐系数是c int64--->8
	其中空结构体size为0
	所以,该结构体的总大小计算方式为: a自己补齐为4,然后与b组8,c自己是8,d再自己补齐为8
	总的就是:4+4+8+8=24
*/
type Part5 struct {
	a bool		//1
	b int32		//4
	c int64		//8
	d struct{}	//0
}

//8+8+8=24
type T1 struct {
	a int8		//1
	b int64		//8
	c int16		//2
}

//8+8=16
type T2 struct {
	a int8		//1
	c int16		//2
	b int64		//8
}

/*
	内存对齐:
		默认系数:在不同平台上的编译器都有自己默认的"对齐系数", 一般而言,32位:4;  64位:8;
		在Go中可以调用unsafe.Alignof来返回相应类型的对齐系数,基本都是2^n,最大也不会超过8.
*/
func main()  {
	t2 := T2{}

	fmt.Println(unsafe.Alignof(t2))		//结构体T2的对齐系数为8
	fmt.Println(unsafe.Sizeof(t2))		//结构体2的大小为24

	/*
		Offsetof返回类型v所代表的结构体字段在结构体中的偏移量，它必须为结构体类型的字段的形式。
	*/
	fmt.Println(unsafe.Offsetof(t2.a))	//0
	fmt.Println(unsafe.Offsetof(t2.c))	//2
	fmt.Println(unsafe.Offsetof(t2.b))	//8
}
