package main

var g *int

//运行: go run -gcflags "-m=2 -1" demo.go
func main()  {
	//escape to heap
	v := 0
	g = &v
}
