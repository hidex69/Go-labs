package task3

import "strings"

func ReverseString(word string) string {
	right := strings.Index(word, ")")

	if right == -1 {
		return word
	}

	left := right

	for word[left] != '(' {
		left--
	}

	w := []rune(word)
	for l, r := left+1, right-1; l < r; l, r = l+1, r-1 {
		w[l], w[r] = w[r], w[l]
	}
	w = append(w[:left], w[left+1:]...)
	w = append(w[:right-1], w[right:]...)

	return ReverseString(string(w))
}
