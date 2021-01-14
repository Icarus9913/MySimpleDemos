package main

import "fmt"

type Valuer interface {
	printf()
}

type Value1 struct {
	num int
}

type Value2 struct {
	num int
}

func (this *Value1) printf() {
	fmt.Println(this.num)
}

func (this *Value2) printf() {
	fmt.Println(this.num)
}

func main() {
	var te1 Valuer
	//te1 = Value1{1}
	te1.printf()

	var te2 Valuer
	te2 = &Value2{2}
	te2.printf()

}
