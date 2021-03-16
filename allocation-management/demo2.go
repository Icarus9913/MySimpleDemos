package main

var p **int

func f1() {
	var x1 *int
	p = &x1

	x2 := x1
	x3 := *p
	x4 := &x3
	_ = x2
	_ = x4
}

func f2() {
	var t **int
	y1 := 1
	y2 := &y1
	y3 := y2

	t = &y2
	p = t
	t = &y3
}

func main() {
	f1()
	f2()
}
