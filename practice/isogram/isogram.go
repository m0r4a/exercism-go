// package isogram provides a function to know if a word is isogram or not
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram recieves a word as a string and returns weather if the word is
// isogram or not in a form of true or false
func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	word = strings.Trim(word, " -")

	for _, letter := range word {
		if !unicode.IsLetter(letter) {
			continue
		}

		if strings.Count(word, string(letter)) > 1 {
			return false
		}
	}
	return true
}
