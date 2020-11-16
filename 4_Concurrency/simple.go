package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Outside a go routine.")
	go func() {
		fmt.Println("Inside a goroutine")
	}()
	fmt.Println("Outside a go routine")
	runtime.Gosched()
}
