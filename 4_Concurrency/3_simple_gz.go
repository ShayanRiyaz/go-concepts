package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	for _, file := range os.Args[1:] { // Collects all list of files passed in on the command line
		compress(file)
	}

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
