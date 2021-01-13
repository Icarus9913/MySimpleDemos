package main

import "fmt"

func main() {
	//fmt.Println("函数a:",a(1))
	//fmt.Println("函数b:",b(1))
	//fmt.Println("函数c:",c(1))
	//fmt.Println("函数d:",d(1))
	//fmt.Println("函数e:",e(1))

	fmt.Println(f(3))
}

func a(n int) int {
	defer fmt.Println("a函数延迟打印:", n)
	n++
	return n
}

func b(n int) int {
	defer func() {
		n += 2
	}()
	return n
}

func c(n int) (i int) {
	defer func() {
		i = 2
	}()
	return n
}

func d(n int) (i int) {
	defer func() {
		i = i + n
	}()
	return n
}

func e(n int) int {
	defer func() {
		fmt.Println("e函数延迟打印", n)
	}()
	n++
	return n
}

/*
注意理解:
	把defer后的当作是一种提前的描述, 该函数中有两个defer,第一个defer func已经描述好了动作.
	而第二个defer f()这里f()是下面的 f = func(){},所以没有描述成功,所以这个defer就没有执行,或者说执行失败
	执行过程: n=3, var f func() 再f = func(){},返回r = n+1就是4, defer f()失败.  继续defer func(){r+=n就是7, 然后将f()失败的recover回来}()
	ps:
	1.defer是在return后执行的.
	2.已经声明好r int为返回值,但是最后return的是n+1,所以最后r就=n+1
*/
func f(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func()
	defer f()

	f = func() {
		r += 2
	}

	return n + 1
}
