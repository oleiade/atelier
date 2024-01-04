package atelier

import (
	"strings"
)

// Trie represents the Trie data structure
type Trie struct {
	Root *TrieNode
}

// NewTrie creates a new Trie
func NewTrie() *Trie {
	return &Trie{Root: NewTrieNode()}
}

// Insert inserts a word into the Trie
func (t *Trie) Insert(word string) {
	currentNode := t.Root

	for _, ch := range word {
		if _, exists := currentNode.Children[ch]; !exists {
			currentNode.Children[ch] = NewTrieNode()
		}
		currentNode = currentNode.Children[ch]
	}

	currentNode.IsEnd = true
}

// Search returns true if the word is in the Trie
func (t *Trie) Search(word string) bool {
	currentNode := t.Root

	for _, ch := range word {
		if node, exists := currentNode.Children[ch]; exists {
			currentNode = node
		} else {
			return false
		}
	}

	return currentNode.IsEnd
}

// StartsWith returns a list of words in the Trie that start with the given prefix.
//
// If no words are found with the given prefix, an empty slice is returned.
func (t *Trie) StartsWith(prefix string) []string {
	currentNode := t.Root
	for _, ch := range prefix {
		if node, exists := currentNode.Children[ch]; exists {
			currentNode = node
		} else {
			return []string{} // No words found with the given prefix
		}
	}
	return t.findWordsFromNode(currentNode, prefix)
}

// EndsWith returns a list of words in the Trie that end with the specified suffix.
//
// It traverses the Trie starting from the root node and collects all the words that match the suffix.
//
// The collected words are returned as a slice of strings.
func (t *Trie) EndsWith(suffix string) []string {
	words := t.findWordsFromNode(t.Root, "")
	var matches []string
	for _, word := range words {
		if strings.HasSuffix(word, suffix) {
			matches = append(matches, word)
		}
	}
	return matches
}

// Autocomplete returns a slice of words in the Trie that have the given prefix
func (t *Trie) Autocomplete(prefix string) []string {
	currentNode := t.Root
	for _, ch := range prefix {
		if node, exists := currentNode.Children[ch]; exists {
			currentNode = node
		} else {
			return []string{} // Prefix not found
		}
	}
	return t.findWordsFromNode(currentNode, prefix)
}

// TrieNode represents each node in the Trie
type TrieNode struct {
	Children map[rune]*TrieNode `json:"children"`
	IsEnd    bool               `json:"is_end"`
}

// NewTrieNode creates a new Trie node
func NewTrieNode() *TrieNode {
	return &TrieNode{Children: make(map[rune]*TrieNode)}
}

// findWordsFromNode performs a DFS from the given node and collects all words
func (t *Trie) findWordsFromNode(node *TrieNode, prefix string) []string {
	var words []string

	if node == nil {
		return words
	}

	if node.IsEnd {
		words = append(words, prefix)
	}

	for ch, childNode := range node.Children {
		words = append(words, t.findWordsFromNode(childNode, prefix+string(ch))...)
	}

	return words
}
