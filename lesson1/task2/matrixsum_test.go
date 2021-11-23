package task2

import "testing"

type testpair struct {
	value  [][]int
	expect int
}

var tests = []testpair{
	{[][]int{{0, 1, 1, 2}, {0, 5, 0, 0}, {2, 0, 3, 3}}, 9},
	{[][]int{{1, 1, 1, 0}, {0, 5, 0, 1}, {2, 1, 3, 10}}, 9},
}

func TestMatrixSum(t *testing.T) {
	for _, pair := range tests {
		value := MatrixSum(pair.value)
		if value != pair.expect {
			t.Error(
				"For", pair.value,
				"expected", pair.expect,
				"got", value,
			)
		}
	}
}
