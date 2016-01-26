package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"
)

var numWorkers int

// Parse and verify command line usage
func init() {
	cpuMultiplier := flag.Int("m", 1, "multiplier for gomaxprocs")

	// File is positional argument so need custom usage message
	flag.Usage = func() {
		fmt.Printf("Usage:\t%s [OPTIONS] <FILE>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	if *cpuMultiplier < 1 {
		fmt.Println("CPU Multiplier cannot be < 1")
		os.Exit(1)
	}

	// Setup for concurrency
	numWorkers = runtime.NumCPU() * *cpuMultiplier
	runtime.GOMAXPROCS(numWorkers)
}

func main() {

	var max int
	var words []string

	fileName := flag.Arg(0)
	start := time.Now()
	trie := NewTrie()
	wordChan := make(chan string)
	resultChan := make(chan *Result, numWorkers)

	// Parse file to build trie
	if err := ParseFile(fileName, trie.Insert); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	// Start the worker processes
	for i := 0; i < numWorkers; i++ {
		go StartWorker(trie, wordChan, resultChan)
	}

	// Parse file again but this time send the words to the workers
	if err := ParseFile(fileName, func(word string) {
		wordChan <- word
	}); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	// Close channel to signal workers to stop waiting for more words
	close(wordChan)

	// Get results from workers and do final check for longest words
	for i := 0; i < numWorkers; i++ {
		r := <-resultChan
		if r.Length > max {
			words = r.Words
			max = r.Length
		} else if r.Length == max {
			words = append(words, r.Words...)
		}
	}

	switch len(words) {
	case 0:
		fmt.Println("No compound word found.")
	case 1:
		fmt.Printf("The longest compound word is: %s (%d)\n", words[0], max)
	default:
		fmt.Printf("Found %d longest compund words of length %d:\n", len(words), max)
		for i := range words {
			fmt.Println("\t", words[i])
		}
	}

	completion := time.Since(start)
	fmt.Println("Time to completion:", completion)
}
