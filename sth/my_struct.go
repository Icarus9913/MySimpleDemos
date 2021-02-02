package main

import (
	"fmt"
	"unsafe"
)

type demo struct {
	c int32
	a struct{}
}

func main()  {

	d := demo{c:1, a:struct{}{}}
	da := d.a
	de := struct {}{}

	fmt.Printf("&(d.a) pointer=%p \nand da pointer=%p \nand de pointer=%p \n",&(d.a),&da,&de)

}



/*
	实际上是未定义行为.是由编译器决定的.虽然实测效果是如此,但是并不能保证以后永远如此,
	编译器完全可能更改字段顺序来节省空间
*/
func TestA() {
	type T1 struct {
		a struct{}
		s string 	//16
		x int64  	//8
		b byte		//1
	}
	t1 := T1{}
	//t1的总: 32 --t1的a: 0 --t1的s: 16 --t1的x: 8 --t1的b: 1
	fmt.Println("t1的总:", unsafe.Sizeof(t1), "--t1的a:", unsafe.Sizeof(t1.a), "--t1的s:", unsafe.Sizeof(t1.s), "--t1的x:", unsafe.Sizeof(t1.x),"--t1的b:", unsafe.Sizeof(t1.b))

	type T2 struct {
		x int64  	//8
		s string 	//16
		b byte		//1
		a struct{}
	}
	t2 := T2{}
	//t2的总: 32 --t2的x: 8 --t2的s: 16 --t2的b: 1 --t2的a: 0
	fmt.Println("t2的总:", unsafe.Sizeof(t2), "--t2的x:", unsafe.Sizeof(t2.x), "--t2的s:", unsafe.Sizeof(t2.s), "--t2的b:", unsafe.Sizeof(t2.b),"--t2的a:", unsafe.Sizeof(t2.a))
}


