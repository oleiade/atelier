package atelier

import (
	"sort"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	t.Parallel()

	trie := NewTrie()

	// Test inserting a single word
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Errorf("Expected 'apple' to be inserted into the trie")
	}

	// Test inserting multiple words
	words := []string{"banana", "cherry", "date"}
	for _, word := range words {
		trie.Insert(word)
		if !trie.Search(word) {
			t.Errorf("Expected '%s' to be inserted into the trie", word)
		}
	}

	// Test inserting duplicate word
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Errorf("Expected 'apple' to be inserted into the trie")
	}
}

func TestTrie_Search(t *testing.T) {
	t.Parallel()

	trie := NewTrie()

	// Test searching for a word that doesn't exist
	if trie.Search("apple") {
		t.Errorf("Expected 'apple' to not be found in the trie")
	}

	// Test searching for a word that exists
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Errorf("Expected 'apple' to be found in the trie")
	}

	// Test searching for a word with a common prefix
	trie.Insert("banana")
	if trie.Search("ban") {
		t.Errorf("Expected 'ban' to not be found in the trie")
	}

	// Test searching for a word with a common prefix
	trie.Insert("cherry")
	if !trie.Search("cherry") {
		t.Errorf("Expected 'cherry' to be found in the trie")
	}
}

func TestTrie_StartsWith(t *testing.T) {
	t.Parallel()

	trie := NewTrie()
	words := []string{"apple", "banana", "cherry", "date"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Test with prefix "a"
	expected := []string{"apple"}
	result := trie.StartsWith("a")
	if !stringSliceEqual(result, expected) {
		t.Errorf("Expected StartsWith result for prefix 'a' to be %v, but got %v", expected, result)
	}

	// Test with prefix "b"
	expected = []string{"banana"}
	result = trie.StartsWith("b")
	if !stringSliceEqual(result, expected) {
		t.Errorf("Expected StartsWith result for prefix 'b' to be %v, but got %v", expected, result)
	}

	// Test with prefix "c"
	expected = []string{"cherry"}
	result = trie.StartsWith("c")
	if !stringSliceEqual(result, expected) {
		t.Errorf("Expected StartsWith result for prefix 'c' to be %v, but got %v", expected, result)
	}

	// Test with prefix "d"
	expected = []string{"date"}
	result = trie.StartsWith("d")
	if !stringSliceEqual(result, expected) {
		t.Errorf("Expected StartsWith result for prefix 'd' to be %v, but got %v", expected, result)
	}

	// Test with prefix "e"
	expected = []string{}
	result = trie.StartsWith("e")
	if !stringSliceEqual(result, expected) {
		t.Errorf("Expected StartsWith result for prefix 'e' to be %v, but got %v", expected, result)
	}
}

func TestTrie_EndsWith(t *testing.T) {
	t.Parallel()

	trie := NewTrie()
	words := []string{"apple", "banana", "cherry", "date"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Test with suffix "e"
	result := trie.EndsWith("e")
	sort.Strings(result)
	expected := []string{"date", "apple"}
	sort.Strings(expected)
	if !stringSliceEqual(result, expected) {
		t.Errorf("Expected EndsWith result for suffix 'e' to be %v, but got %v", expected, result)
	}

	// Test with suffix "a"
	expected = []string{"banana"}
	result = trie.EndsWith("a")
	if !stringSliceEqual(result, expected) {
		t.Errorf("Expected EndsWith result for suffix 'a' to be %v, but got %v", expected, result)
	}

	// Test with suffix "y"
	expected = []string{"cherry"}
	result = trie.EndsWith("y")
	if !stringSliceEqual(result, expected) {
		t.Errorf("Expected EndsWith result for suffix 'y' to be %v, but got %v", expected, result)
	}

	// Test with suffix "x"
	expected = []string{}
	result = trie.EndsWith("x")
	if !stringSliceEqual(result, expected) {
		t.Errorf("Expected EndsWith result for suffix 'x' to be %v, but got %v", expected, result)
	}
}

func TestTrie_Autocomplete(t *testing.T) {
	t.Parallel()

	trie := NewTrie()

	// Insert words into the trie
	words := []string{"apple", "banana", "cherry", "date"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Test autocomplete with prefix "a"
	expected := []string{"apple"}
	result := trie.Autocomplete("a")
	if !stringSliceEqual(result, expected) {
		t.Errorf("Expected autocomplete result for prefix 'a' to be %v, but got %v", expected, result)
	}

	// Test autocomplete with prefix "b"
	expected = []string{"banana"}
	result = trie.Autocomplete("b")
	if !stringSliceEqual(result, expected) {
		t.Errorf("Expected autocomplete result for prefix 'b' to be %v, but got %v", expected, result)
	}

	// Test autocomplete with prefix "c"
	expected = []string{"cherry"}
	result = trie.Autocomplete("c")
	if !stringSliceEqual(result, expected) {
		t.Errorf("Expected autocomplete result for prefix 'c' to be %v, but got %v", expected, result)
	}

	// Test autocomplete with prefix "d"
	expected = []string{"date"}
	result = trie.Autocomplete("d")
	if !stringSliceEqual(result, expected) {
		t.Errorf("Expected autocomplete result for prefix 'd' to be %v, but got %v", expected, result)
	}

	// Test autocomplete with prefix "e"
	expected = []string{}
	result = trie.Autocomplete("e")
	if !stringSliceEqual(result, expected) {
		t.Errorf("Expected autocomplete result for prefix 'e' to be %v, but got %v", expected, result)
	}
}

func stringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestTrie_FindWordsFromNode(t *testing.T) {
	t.Parallel()

	trie := NewTrie()
	words := []string{"apple", "banana", "cherry", "date"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Test finding words from a node with prefix "a"
	node := trie.Root.Children['a']
	expected := []string{"apple"}
	result := trie.findWordsFromNode(node, "a")
	if !stringSliceEqual(result, expected) {
		t.Errorf("Expected FindWordsFromNode result for prefix 'a' to be %v, but got %v", expected, result)
	}

	// Test finding words from a node with prefix "b"
	node = trie.Root.Children['b']
	expected = []string{"banana"}
	result = trie.findWordsFromNode(node, "b")
	if !stringSliceEqual(result, expected) {
		t.Errorf("Expected FindWordsFromNode result for prefix 'b' to be %v, but got %v", expected, result)
	}

	// Test finding words from a node with prefix "c"
	node = trie.Root.Children['c']
	expected = []string{"cherry"}
	result = trie.findWordsFromNode(node, "c")
	if !stringSliceEqual(result, expected) {
		t.Errorf("Expected FindWordsFromNode result for prefix 'c' to be %v, but got %v", expected, result)
	}

	// Test finding words from a node with prefix "d"
	node = trie.Root.Children['d']
	expected = []string{"date"}
	result = trie.findWordsFromNode(node, "d")
	if !stringSliceEqual(result, expected) {
		t.Errorf("Expected FindWordsFromNode result for prefix 'd' to be %v, but got %v", expected, result)
	}

	// Test finding words from a node with prefix "e"
	node = trie.Root.Children['e']
	expected = []string{}
	result = trie.findWordsFromNode(node, "e")
	if !stringSliceEqual(result, expected) {
		t.Errorf("Expected FindWordsFromNode result for prefix 'e' to be %v, but got %v", expected, result)
	}
}
