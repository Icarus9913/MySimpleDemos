package main

import (
	"fmt"
	"time"
)

var TimeNowStr string = "2006-01-02 15:04:05"

func init()  {
	go func() {
		for {
			tim_t := time.Now()
			TimeNowStr = tim_t.Format("2006-01-02 15:04:05")
			time.Sleep(time.Second)
		}
	}()
}

func main()  {
	log("my stupid test = %v","bad")
	
}

func log(format string, v ...interface{}) {
	fmt.Println(fmt.Sprintf("[debug][%s]", TimeNowStr), fmt.Sprintf(format, v...))
}