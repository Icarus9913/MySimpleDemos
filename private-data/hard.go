package main

import (
	"fmt"
	"unsafe"
)

//4+4+8=16
type tt struct {
	a int8  //1
	b int8  //1
	c int32 //4
	d int64 //8
}

/*
	16个字节分4块,每1块占用4个字节,而1个字节是由两位16进制数组成.
	-->1Byte=8bit, xxxx,xxxx表示0到255,而1个16进制数最多表示0-15,而要表示到255,则需要两个16进制数FF(15*16+15)
*/

//01,02,00,00|03,00,00,00|04,00,00,00|00,00,00,00
func newTT() *tt {
	return &tt{
		a: 1,
		b: 2,
		c: 3,
		d: 4,
	}
}

func main() {
	t := newTT()
	p := unsafe.Pointer(t)

	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	m := func() unsafe.Pointer {
		temp := uintptr(p)
		return unsafe.Pointer(temp)
	}()
	i := (*int)(m)		//int在64位机器中,占8个字节,所以此时的i拿到的是01,02,00,00|03,00,00,00  -->操作系统的大端与小端-> 03,00,00,02,01->这是16进制,然后转成10进制为12884902401
	fmt.Println(*i)

	//----------------------------
	m1 := func() unsafe.Pointer {
		temp := uintptr(p)
		return unsafe.Pointer(temp)
	}()
	i1 := (*int8)(m1)
	fmt.Println(*i1)	//1

	//----------------------------
	m2 := func() unsafe.Pointer {
		temp := uintptr(p)
		temp += unsafe.Sizeof(int8(0))
		return unsafe.Pointer(temp)
	}()
	i2 := (*int8)(m2)
	fmt.Println(*i2)	//2

	//==============================
	m3 := func() unsafe.Pointer {
		temp := uintptr(p)
		temp += unsafe.Sizeof(int32(0))
		return unsafe.Pointer(temp)
	}()
	i3 := (*int32)(m3)
	fmt.Println(*i3)	//3

	//-----------------------------
	m4 := func() unsafe.Pointer {
		temp := uintptr(p)
		temp += unsafe.Sizeof(int64(0))
		return unsafe.Pointer(temp)
	}()
	i4 := (*int64)(m4)
	fmt.Println(*i4)	//4
}

/*
	大端模式，是指数据的低位保存在内存的高地址中，而数据的高位，保存在内存的低地址中；
	小端模式，是指数据的低位保存在内存的低地址中，而数据的高位保存在内存的高地址中。（大部分的操作系统都是小端，而通讯协议是大端）
*/
func IsLittleEndian() bool {
	var i int32 = 0x01020304

	// 下面这两句是为了将int32类型的指针转换为byte类型的指针
	u := unsafe.Pointer(&i)
	pb := (*byte)(u)

	b := *pb // 取得pb位置对应的值

	// 由于b是byte类型的,最多保存8位,那么只能取得开始的8位
	// 小端: 04 (03 02 01)
	// 大端: 01 (02 03 04)
	return (b == 0x04)	//true表示是小端,因为计算十六进制是从右向左计算,所以04是数据小位，如果存放地址，内存地址低位是04，则表明小位对小位.
}