package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func timerStart(c chan bool) {
	count := 5
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Second)
	}
	c <- true
}

func getUserClose(ch chan os.Signal, flag chan bool) {
	<-ch
	fmt.Println("Stop by user")
	close(flag)
}

func main() {
	interrupt := make(chan os.Signal, 1)
	c := make(chan bool)
	signal.Notify(interrupt, os.Interrupt)
	go getUserClose(interrupt, c)
	fmt.Println("Hello World")
	go timerStart(c)

	switch <-c {
	case true:
		fmt.Println("Goodbye world")
	default :
		fmt.Println(" ")
	}
}
