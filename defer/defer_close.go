package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//Correct a mistake about defer in the code below
func main()  {
	f,err := os.Open("file")
	defer f.Close()

	if nil!=err{
		fmt.Println("失败")
		return
	}

	b,err := ioutil.ReadAll(f)
	println(string(b))
}


//my fucking change
func init()  {
	f, err := os.Open("file")
	if nil!=err{
		return
	}
	defer f.Close()
}


/*
	explain:
	假如os.open()中出现故障，有两种情况，一申请出file，但其不可用；另一种则是未申请到file，通常file获得的返回值是nil
	当然，它们的err都不会是nil
	如果是先执行(声明)defer xx.Close()的话，在第二种情况下，可能会在conn为nil的时候依然执行了conn.Close()，假如这个Close()函数保护不够充分，则容易会产生panic

故建议：
	先对err进行判断，如果err不是nil，那么肯定已经有错误产生了，先处理错误
	若err是nil，那么conn也是可用的，这时候声明defer conn.Close()便不会有什么问题了
*/