package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup // A wait group doesn't need to be initialized
	var i int = -1        // Because we want to reference i outside the for loop we declare variables there
	var file string
	for i, file = range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			compress(filename)
			wg.Done()
		}(file)
	}
	wg.Wait()

	fmt.Printf("Compressed %d files\n", i+1)
}
func compress(filename string) error {
	in, err := os.Open(filename) // Opens the source file for reading
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz") // Opens a destination file,
	// with the .gz extension added to
	if err != nil { // the source file's name
		return err
	}
	defer out.Close()

	gzout := gzip.NewWriter(out) // the gzip. Writer compresses data and then writes it
	// to the underlying file.
	_, err = io.Copy(gzout, in) // The io.Copy functin does all the copying for you
	gzout.Close()

	return err
}
