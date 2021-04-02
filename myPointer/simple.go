package main

import "fmt"

/*
	由C语言中的二级指针引出的问题.当在主函数中定义一个指针变量,然后传给子函数去malloc,就会碰到坑,此时需要用到二级指针
	注意: *p是指针变量p指向地址的值(实体的值)； p是实体的地址；  &p是指针变量p的地址
*/

type person struct {
	age int
}
func main()  {
	p := &person{}		//p是一个指针变量, p的值是Person{}的地址
	funcChangeP(p)		//实参,传入一个p,其中p的值是Person{}的地址, bad
	p.ChangeP()			//bad
	p.ChangeAge()		//good
	fmt.Println(p)
}

func funcChangeP(p *person) { 	//行参, copy了指针变量,只不过p的值还是实参Person{}的地址
	pp := person{20}
	p = &pp 					//改变p的值,但是此时p的地址是该函数,copy出来的一个,而不是主函数中的指针变量p
}

func (p *person) ChangeP() { 	//这里也是copy了指针变量出来,p的值还是实参Person{}的地址, 这里变量p的地址与主函数中p的地址不一样,因为是copy
	pp := person{20}
	p = &pp						//改变p的值,只是在该方法里,改变了这个方法里copy出来的p变量的值,原来主函数中的p的值未改变
}

func (p *person) ChangeAge() {	//这里改变的是copy变量指向的实体的值. 注意此处改变的是实体.
	p.age = 20					//所以不管是copy的p还是主函数中实际的p,他们指向的实体的值被改了.
}



