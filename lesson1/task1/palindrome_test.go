package task1

import "testing"

type testpair struct {
	value  string
	expect bool
}

var tests = []testpair{
	{"owo", true},
	{"a", true},
	{"awo", false},
	{"aaavaaa", true},
}

func TestIsPolindrome(t *testing.T) {
	for _, pair := range tests {
		value := IsPalindrome(pair.value)
		if value != pair.expect {
			t.Error(
				"For", pair.value,
				"expected", pair.expect,
				"got", value,
			)
		}
	}
}
