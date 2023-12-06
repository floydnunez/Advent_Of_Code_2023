package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGKILL)

	go func() {
		for {
			<-sigChan
			fmt.Println("Received SIGHUP, printing stack trace:")
			printStackTrace()
		}
	}()
	var arr []byte

	for {
		arr = append(arr, make([]byte, 100000*1024*1024)...)
	}
}

func printStackTrace() {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			fmt.Print(string(buf[:n]))
			break
		}
		buf = make([]byte, 2*len(buf))
	}
}
