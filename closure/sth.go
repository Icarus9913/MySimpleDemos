package main

import "fmt"

//详情请看closure_alloc这张图

func main() {
	fs := crete()
	for i := 0; i < len(fs); i++ {
		fs[i]()
	}

	fmt.Println("============")

	fs = bury()
	for i := 0; i < len(fs); i++ {
		fs[i]()
	}

}

/*
	for循环中的变量i逃逸到堆上,而fs[i]里面的print都是指向这个逃逸变量i的地址,
	所以最后执行全部打印的是2,因为for循环后i的值变成了2
*/

func crete() (fs [2]func()) {
	for i := 0; i < 2; i++ {
		fs[i] = func() {
			fmt.Println(i)
		}
	}
	return
}

/*
	这里的每个j都压栈了,都是独立的j
*/
func bury() (fs [2]func()) {
	for i := 0; i < 2; i++ {
		func(j int) {
			fs[j] = func() {
				fmt.Println(j)
			}
		}(i)
	}
	return
}

