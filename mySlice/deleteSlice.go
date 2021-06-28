package main

import "fmt"

func main()  {
	demo2()
	
}

// 删除4
func demo1()  {
	s := []int{1,2,3,4,5,6,7}
	val := 4
	for i,v := range s{
		if v==val{
			s = append(s[:i],s[i+1:]...)
		}
	}
	fmt.Println(s)
}

// 清空切片
func demo2()  {
	s := []int{1,2,3,4}
	fmt.Println(s[:0])
	fmt.Println(s[len(s):])		//s[4:] 不会越界;  直接访问s[4]会越界panic

	s = append(s[:0],s[len(s):]...)
	fmt.Println(s,len(s),cap(s))
}
