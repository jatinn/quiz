// Was not sure if we could use external libraries so implemented smallest trie
// needed to solve the problem.

package main

import "bytes"

type node struct {
	character rune
	children  map[rune]*node
	isEOW     bool
}

type Trie struct {
	root *node
}

func NewTrie() *Trie {
	return &Trie{
		&node{
			children: make(map[rune]*node),
		},
	}
}

func (t *Trie) Insert(word string) {
	currentNode := t.root
	for _, r := range word {
		if _, found := currentNode.children[r]; !found {
			currentNode.children[r] = &node{
				character: r,
				children:  make(map[rune]*node),
			}
		}
		currentNode = currentNode.children[r]
	}
	currentNode.isEOW = true
}

// For a given word return all substrings it comprises of if it is a compound
// word otherwise returns an empty list.
func (t *Trie) FindParts(word string, isSuffix bool) (parts []string) {
	var pBuffer bytes.Buffer
	var prefix string
	currentNode := t.root

	// Iterate over each character of the word and check whether it is a terminal
	// character in the trie. If it is recursively call on the remaining suffix
	// of the word. If suffix does not yield a compound word continue to grow
	// prefix and keep checking suffixes until either a match is found or the
	// word itself is found in the trie.
	for _, r := range word {
		if _, found := currentNode.children[r]; !found {
			return
		}
		currentNode = currentNode.children[r]
		pBuffer.WriteRune(r)
		if currentNode.isEOW {
			prefix = pBuffer.String()
			if !isSuffix && word == prefix {
				return
			} else if isSuffix && word == prefix {
				parts = append(parts, prefix)
			}
			sParts := t.FindParts(word[pBuffer.Len():], true)
			if len(sParts) > 0 {
				parts = append(parts, prefix)
				parts = append(parts, sParts...)
			}
		}
		// Compound words could be formed with multiple different sets of
		// sub words. Break once we have found any set.
		if len(parts) > 0 {
			return
		}
	}
	return
}

func (t *Trie) IsCompound(word string) bool {
	return len(t.FindParts(word, false)) > 0
}
