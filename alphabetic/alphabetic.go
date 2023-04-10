package alphabetic

import (
	"fmt"
	"strconv"
	"strings"
)

var Alphabet []string = alphabetGenerator()

func alphabetGenerator() []string {
	var alphabet []string

	for ch := 'a'; ch <= 'z'; ch++ {
		var char string = strconv.QuoteRune(ch)
		char = strings.ReplaceAll(char, "'", "")
		alphabet = append(alphabet, char)
	}
	for i := 0; i <= 9; i++ {
		alphabet = append(alphabet, fmt.Sprint(i))
	}

	return alphabet
}
