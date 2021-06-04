package main

import (
	"fmt"
	"sync"
)

func main() {
	var o sync.Once
	oneBody := func() {
		fmt.Println("Ha Ha")
	}

	busBody := func() {
		fmt.Println("Ga Ga")
	}

	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			o.Do(oneBody)
			busBody()
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}
