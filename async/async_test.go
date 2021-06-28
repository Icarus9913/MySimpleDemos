package async

import (
	"fmt"
	"testing"
	"time"
)

func DoneAsync() int {
	fmt.Println("Warming up...")
	time.Sleep(3*time.Second)
	fmt.Println("Done...")
	return 1
}

func TestExec(t *testing.T) {
	fmt.Println("Let's start...")
	future := Exec(func() interface{} {
		return DoneAsync()
	})
	fmt.Println("Done is running...")
	val := future.Await()
	fmt.Println(val)
}
