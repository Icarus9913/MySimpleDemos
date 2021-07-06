package main

import "fmt"

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
	若一开始a的len置为0的话,最后出现的情况是cap3里都有值,但却因为len是0,而打印什么都没有
*/
