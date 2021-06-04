package main

import (
	"fmt"
	"sync"
)

func main() {
	langCap := sync.Map{}
	langCap.Store("CN", "Beijing")
	langCap.Store("US", "NewYork")
	langCap.Store("JP", "Tokyo")

	//for key, value := range langCap {
	//	fmt.Println(key, value)
	//}

	langCap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		if key == "US" {
			return false
		}
		return true
	})


	var msg interface{}
	msg = "123"
	var msg2 int32
	//msg2 = "123"
	fmt.Println(msg, msg2)
}
