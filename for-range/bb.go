package main

import "fmt"

var slice []func()

func main() {
	sli := []int{1, 2, 3, 4, 5}

	/*	for _,v := range sli{
		fmt.Println(v,":",&v)
	}*/

	for _, v := range sli {
		//temp := v
		slice = append(slice, func() {
			//fmt.Println(temp * temp)
			fmt.Println(v * v)
		})
	}

	for _, val := range slice {
		val()
	}
}

//Go 陷阱之 for 循环迭代变量
//https://www.jianshu.com/p/0e2e353b6c7d