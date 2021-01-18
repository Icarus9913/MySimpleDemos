package main

import "log"

//What will be printed when the code below is executed?
//What will be the exit code after the code below is executed?
func ff() {
	defer func() {
		if r := recover(); nil != r {
			log.Printf("recover:%#v", r)
		}
	}()
	panic(1)
	panic(2)
}

func main() {
	ff()
}
