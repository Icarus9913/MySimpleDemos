package main

import "fmt"

const N = 3

func main()  {
	m := make(map[int]*int)

	for i:=0;i<N;i++{
		j := i
		m[i]= &j

	}

	for _,v := range m{
		fmt.Println("*v的值:",v)
	}

}
