package atelier

import (
	"reflect"
)

// MapAddressToWord maps a given value's address to a word from the WordList.
//
// This function proves useful when debugging as it allows you to
// identify a variable holding a pointer by a unique name, as opposed
// to having to look at the address.
//
// It uses a hash function to calculate the index in the WordList based on the address.
// The value must be a pointer type.
//
// The WordList is a global variable containing a list of words.
//
// Returns the word corresponding to the address.
func MapAddressToWord(value any) string {
	// hashAddress takes a pointer and hashes it to an integer.
	hashAddressFunc := func(address uintptr) int {
		return int(address) % len(WordList)
	}

	valuePointer := reflect.ValueOf(value).Pointer()
	hash := hashAddressFunc(valuePointer)
	return WordList[hash]
}
