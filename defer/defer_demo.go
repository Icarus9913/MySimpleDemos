package main

import "fmt"

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

func DeferFunc4() (t int) {
	defer func(i int) {
		fmt.Println(i)
		fmt.Println(t)
	}(t)
	t = 1
	return 2
}

func main() {
	fmt.Println(DeferFunc1(1))
	fmt.Println(DeferFunc2(1))
	fmt.Println(DeferFunc3(1))
	fmt.Println(DeferFunc4())
}

/*
	DeferFunc1:
	1.将返回值t赋值为传入的i，此时t为1
	2.执行return语句将t赋值给t（等于啥也没做）
	3.执行defer方法，将t + 3 = 4
	4.函数返回 4, 因为t的作用域为整个函数所以修改有效。
*/

/*
	DeferFunc2:
	1.创建变量t并赋值为1
	2.执行return语句，注意这里是将t赋值给返回值，此时返回值为1（这个返回值并不是t）
	3.执行defer方法，将t + 3 = 4
	4.函数返回返回值1
也可按照以下代码理解：
func DeferFunc2(i int) (result int) {
    t := i
    defer func() {
        t += 3
    }()
    return t
}
上面的代码return的时候相当于将t赋值给了result，当defer修改了t的值之后，对result是不会造成影响的。
*/

/*
	DeferFunc3:
	1.首先执行return将返回值t赋值为2
	2.执行defer方法将t + 1
	3.最后返回 3
*/

/*
	DeferFunc4:
	1.初始化返回值t为零值 0
	2.首先执行defer的第一步，赋值defer中的func入参t为0
	3.执行defer的第二步，将defer压栈
	4.将t赋值为1
	5.执行return语句，将返回值t赋值为2
	6.执行defer的第三步，出栈并执行,因为在入栈时defer执行的func的入参已经赋值了，
	此时它作为的是一个形式参数，所以打印为0；相对应的因为最后已经将t的值修改为2，所以再打印一个2
*/
