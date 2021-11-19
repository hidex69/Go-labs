package main

import (
	"fmt"
	"strings"
)

func main() {
	//task1
	fmt.Println(isPalindrome("a"))

	//task2
	m := [][]int{{1, 1, 1, 0}, {0, 5, 0, 1}, {2, 1, 3, 10}}
	fmt.Println(matrixSum(m))

	//task3
	fmt.Println(reverseString("foo(bar(baz))blim"))
}

func isPalindrome(word string) bool {
	result := true
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		if word[i] != word[j] {
			result = false
		}
	}
	return result
}

func matrixSum(matrix [][]int) int {
	sum := 0

	for j := 0; j < len(matrix[0]); j++ {
		for i := 0; i < len(matrix); i++ {
			if matrix[i][j] == 0 {
				break
			} else {
				sum += matrix[i][j]
			}
		}
	}

	return sum
}

func reverseString(word string) string {
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

	return reverseString(string(w))
}
