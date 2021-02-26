package main

import "fmt"

func createFunc(v int) func() {
	return func() {
		fmt.Println("传参数给函数打印的结果，值为：", v)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	var funcArr []func()

	for _, val := range arr {
		// 直接打印出结果
		fmt.Println("直接打印的结果，值为：", val)

		// val的生命周期是整个for语句，它会依次被赋值为arr中的每一个元素值，也就是整个for语句中，val只有一份
		// 每次循环时，会把func(){}的定义连同val打包成了一个闭包函数放到了funcArr中，要注意的是，这时并没有去执行val
		// 执行的时间点是在后面的for...range funcArr{}语句，而这时，for循环语句早已经执行完，val的值也就是最后的那个元素值了
		funcArr = append(funcArr, func() {
			fmt.Println("直接给函数打印的结果，值为：", val)
		})

		// 每次循环都会把val的值拷贝一份给copyVal，然后把func(){}连同copyVal打包成一个函数放入funcArr中
		// 注意这里是:=，每次都会重新声明，它的作用域是for...range arr{}内，也就是会有5个copyVal对应着不同的val值
		// 最后执行时，每次打印的是当时打包进func(){}里的那个copyVal，因此值是对的。
		copyVal := val
		funcArr = append(funcArr, func() {
			fmt.Println("声明一个变量再给函数打印的结果，值为：", copyVal)
		})

		// 给函数传参，每次都会先拷贝一份再给函数
		funcArr = append(funcArr, createFunc(val))
	}

	// 依次执行
	for _, f := range funcArr {
		f()
	}
}