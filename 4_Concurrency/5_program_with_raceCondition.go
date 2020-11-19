package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup // We'll use a wait group to monitor a group of goroutines
	w := newWords()

	for _, f := range os.Args[1:] { // The main loop uses the pattern used previously
		wg.Add(1)
		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
	fmt.Println("Worlds that appear more than once:")
	for word, count := range w.found {
		if count > 1 { // At the end of the program, we print the found word
			fmt.Printf("%s: %d\n", word, count)
		}
	}
}

type words struct { // We track words in a struct. We could use a map tyle, but using a struct here makes the next refactor easier.
	found map[string]int
}

func newWords() *words { // Creates a new word instance
	return &words{found: map[string]int{}}
}

func (w *words) add(word string, n int) { // Tracks the number of times we've seen this word
	count, ok := w.found[word]
	if !ok { // If the word isn't already tracked, add it. Otherwise, increment the count.
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}

func tallyWords(filename string, dict *words) error { // Open a file, parse it's contents
	file, err := os.Open(filename) // and count the words that appear.
	if err != nil {                // Copy function does all the copying for us
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}
