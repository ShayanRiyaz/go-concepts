package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	go echo(os.Stdin, os.Stdout) // Calls the function echo as a goroutine
	time.Sleep(30 * time.Second) // Sleeps for 30 seconds
	fmt.Println("Timed out.")    // Prints out a message saying we're done sleeping
	os.Exit(0)                   // Exits the program. This stops the go routine
}

func echo(in io.Reader, out io.Writer) { // The echo function is a normal function
	io.Copy(out, in) // io.Copy copies data to an os.Writer from an os.Reader
}
