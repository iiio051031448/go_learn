package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("nuclear launch detected.")
}

func commencingCountDown(canLaunch chan int) {
	c := time.Tick(1 * time.Second)
	for countDown := 20; countDown > 0; countDown-- {
		fmt.Println(countDown)
		<-c
	}
	canLaunch <- -1
}

func isAbort(abort chan int) {
	os.Stdin.Read(make([]byte, 1))
	abort <- -1
}

func main() {
	fmt.Println("commencing countdown")

	abort := make(chan int)
	canLaunch := make(chan int)
	go isAbort(abort)
	go commencingCountDown(canLaunch)
	select {
	case <-abort:
	case <-canLaunch:
		fmt.Println("launch aborted")
		return
	}
	launch()
}