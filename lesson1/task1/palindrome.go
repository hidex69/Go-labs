package task1

func IsPalindrome(word string) bool {
	result := true
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		if word[i] != word[j] {
			result = false
		}
	}
	return result
}
