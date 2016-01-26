package main

import (
	"bufio"
	"os"
	"strings"
	"unicode/utf8"
)

type Result struct {
	Words  []string
	Length int
}

func StartWorker(t *Trie, words <-chan string, result chan<- *Result) {
	var longestWords []string
	var max, length int

	// Check incomming words and find the longest compound words and save
	// in local cache until chan is closed
	for word := range words {
		length = utf8.RuneCountInString(word)
		if length >= max && t.IsCompound(word) {
			if length > max {
				longestWords = []string{word}
				max = length
			} else {
				longestWords = append(longestWords, word)
			}
		}
	}

	// Send back local results to main thead
	result <- &Result{
		Words:  longestWords,
		Length: max,
	}
}

// Parse file and applyi the given function to each line
func ParseFile(name string, fn func(word string)) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	var word string
	for s.Scan() {
		// Normalize words to lowercase and only apply function to non empty words
		word = strings.ToLower(s.Text())
		if word != "" {
			fn(word)
		}
	}
	return s.Err()
}
