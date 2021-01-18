package main

import (
	"fmt"
	"sort"
)

//Add code to line A to sort s in ascending order
type SSS struct {
	v int
}

func main()  {
	s := []SSS{{1},{3},{5},{2}}
	//A
	fmt.Printf("%#v",s)
}

//my fucking change
func idiot()  {
	s := []SSS{{1},{3},{5},{2}}

	sort.Slice(s, func(i, j int) bool {
		return s[i].v < s[j].v
	})
	fmt.Printf("%#v",s)
}