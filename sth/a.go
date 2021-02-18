package main

import (
	"fmt"
)

func main()  {
	a := 15
	fmt.Println("祖先:",a)
	//Re(&a)
	swap(&a)
	fmt.Println("原来:",a)
}

func Re(a *int) {
	fmt.Println("第一次",a)
	cp := *a	//cp的值为a的值
	a = &cp		//a的地址等于cp的地址
	*a = 11		//a的值为11
	fmt.Println("返回",a)
}

func swap(a *int)  {
	*a =5

}
