package treeHeight

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	input    []int
	n        int
	expected int
}

func TestTreeHeight(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    []int{9, 7, 5, 5, 2, 9, 9, 9, 2, -1},
			n:        10,
			expected: 4,
		},
		{
			input:    []int{4, -1, 4, 1, 1},
			n:        5,
			expected: 3,
		},
		{
			input:    []int{-1, 0, 4, 0, 3},
			n:        5,
			expected: 4,
		},
	} {
		result := TreeHeight(tst.n, tst.input)
		require.Equal(t, tst.expected, result)
	}
}
