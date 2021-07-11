package main

import "fmt"

/*
	若一开始a的len置为0的话,最后出现的情况是cap3里都有值,但却因为len是0,而打印什么都没有.
	解析:
		起初a的len与cap都是3; b的len是0,cap是3.
		第8行,append操作先将数据存入b指针实体中, 然后再去改变b的len值.及先执行append右边,最后执行左边的b=
		19行中,相当于b的底层指针数据成功append进去了,且a的len是3,所以a指向的实体能成功显示出来.
				对于b而言,虽然指针实体的数据成功变了,但是他的len还是2,没能改变.
*/
func main() {
	a := make([]int, 3) //注意此处a的len与cap都是3
	b := a[:0]          //b的len是0,cap是3, 此刻b指针与a的指针指向同一个地方
	b = append(b, 1, 2)

	c := []int{3}
	_ = append(b, c...) //到底发生了啥? 为什么c的值能赋值给a[2]?
	fmt.Println(b)      // [1,2]  ->b没能追加c
	fmt.Println(a)      // [1,2,3]	-> a成功追加c
}


/*
	看go/src/runtime/slice.go -->growslice
	append后的cap首先是9,走case et.size == sys.PtrSize的roundupsize函数, 注意:sys.PtrSize应该是8个字节的意思
	9*8=72, divRoundUp得到9, size_to_class8[9]=7, class_to_size[7]=80, 最后80/8等于10字节
*/
func another()  {
	arr1 := [4]int{1,2,3,4}
	slice1 := arr1[:3]
	fmt.Println(slice1,len(slice1),cap(slice1))			//[1 2 3] 3 4
	slice1 = append(slice1,5000,6000,7000,8000,9000,10000)
	fmt.Println(slice1,len(slice1),cap(slice1))			//[1 2 3 5000 6000 7000 8000 9000 10000] 9 10
}

/*
	v刚开始的len是3,cap是5
	append先扩v的len,并把6这个值赋值给实体len4位置,因此替换了.
	append控制len!!!!
*/
func sth()  {
	arr := []int{1, 2, 3, 4, 5}
	v := arr[:3]
	fmt.Println(v, len(v), cap(v)) // [1 2 3] 3 5
	v = append(v, 6)
	fmt.Println(v, len(v), cap(v)) // [1 2 3 6] 4 5
}
