package task3

import "testing"

type testpair struct {
	value  string
	expect string
}

var tests = []testpair{
	{"(bar)", "rab"},
	{"foo(bar)baz", "foorabbaz"},
	{"foo(bar)baz(blim)", "foorabbazmilb"},
	{"foo(bar(baz))blim", "foobazrabblim"},
}

func TestReverseString(t *testing.T) {
	for _, pair := range tests {
		value := ReverseString(pair.value)
		if value != pair.expect {
			t.Error(
				"For", pair.value,
				"expected", pair.expect,
				"got", value,
			)
		}
	}
}
