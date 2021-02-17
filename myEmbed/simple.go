package main

import (
	_ "embed"
	"fmt"
)

/*
	如果abc.txt文件在某个文件夹下，例如a/abc.txt就直接写成a/abc.txt即可
*/

//go:embed abc.txt
var str string

func main()  {
	fmt.Println(str)
}
