package main

import (
	"time"
	"fmt"
)

//每秒打印hello world！
func main() {
	fmt.Println("hello world!")
	for {
		fmt.Println("sleep 1s")
		time.Sleep(time.Second)
	}
}
