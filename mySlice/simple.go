package main

import "fmt"

/*
slice的底层结构
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
*/

/*
	首先bad()使用直接切片，此时nums的len与cap都是3，
	执行badAppendNum后，因为nums的cap是3都塞满了值，需要扩容，然后v指向了新的地址空间，并且把nums之前的值copy出来了
	所以v的值发生改变了，注意这里v的值是底层数组的头节点地址.结论是：v与nums所指向的值发生了改变。
*/
func bad() {
	nums := []int{1, 2, 3}
	badAppendNum(nums)
	fmt.Println("bad外：", nums)
}

func badAppendNum(v []int) {
	v = append(v, 4)
	fmt.Println("bad里：", v)
}

/*
	对原nums进行直接操作，而没有copy nums
*/
func good() {
	nums := []int{1, 2, 3}
	goodAppendNum(&nums)
	fmt.Println("good外：", nums)
}

func goodAppendNum(v *[]int) {
	*v = append(*v, 4)
	fmt.Println("good里:", *v)
}

/*
	这里先初始化一个len为3，cap为4的slice，然后直接copy这个slice去append，
	可以发现在主函数中虽然直接打印v是没有变化的，但是打印v[:4]后会发现变了，
	因为我们直接操作修改了slice底层的array指针数组，主函数直接打印v没有变化是因为它此时的len还是3没有改变
	而cap是4，说明还有一个空间是可以去append，而不用扩容重新指向新的slice
*/
func something() {
	v := make([]int, 3, 4)
	v[0] = 1
	v[1] = 2
	v[2] = 3
	fmt.Println(v, len(v), cap(v))
	fmt.Println(v[:4])

	change(v)
	fmt.Println(v, len(v), cap(v))
	fmt.Println(v[:4])
}

func change(v []int) {
	v = append(v, 4)
	fmt.Println(v, len(v), cap(v))
}

func main() {
	something()
}

func interesting() {
	a := []int{1, 2, 3, 4, 5}
	b := a[:2]   //表示[1,2],且len为2,cap是5
	c := a[:2:3] //表示[1,2],且len为2,cap为3
	fmt.Println(b, len(b), cap(b))
	fmt.Println(c, len(c), cap(c))
}
