package main
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	signalChan := make(chan os.Signal,1)
	signal.Notify(signalChan,
		syscall.SIGINT,			//Ctrl+c
		syscall.SIGTSTP,		//Ctrl+z
		syscall.SIGQUIT)		//Ctrl+\
	s:=<-signalChan
	fmt.Println("\nget signal:",s)
}
