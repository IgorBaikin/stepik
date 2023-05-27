package networkPack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	arrival  []int
	duration []int
	bufSize  int
	expected []int
}

func TestNetworkPackProc(t *testing.T) {
	for _, tst := range [...]test{
		{
			arrival:  []int{},
			duration: []int{},
			bufSize:  1,
			expected: []int{},
		},
		{
			arrival:  []int{0},
			duration: []int{0},
			bufSize:  1,
			expected: []int{0},
		},
		{
			arrival:  []int{0, 0},
			duration: []int{1, 1},
			bufSize:  1,
			expected: []int{0, -1},
		},
		{
			arrival:  []int{0, 1},
			duration: []int{1, 1},
			bufSize:  1,
			expected: []int{0, 1},
		},
		{
			arrival:  []int{2, 4, 10, 15, 19},
			duration: []int{9, 8, 9, 2, 1},
			bufSize:  2,
			expected: []int{2, 11, -1, 19, 21},
		},
		{
			arrival:  []int{0, 0, 0, 1, 1, 1, 1, 1},
			duration: []int{0, 0, 0, 0, 0, 1, 2, 3},
			bufSize:  2,
			expected: []int{0, 0, 0, 1, 1, 1, 2, -1},
		},
		{
			arrival:  []int{0, 0, 0, 1, 1, 1, 1, 1},
			duration: []int{0, 0, 0, 1, 0, 0, 2, 3},
			bufSize:  2,
			expected: []int{0, 0, 0, 1, 2, -1, -1, -1},
		},
		{
			arrival:  []int{999999, 1000000, 1000000, 1000000, 1000000},
			duration: []int{1, 0, 1, 0, 0},
			bufSize:  1,
			expected: []int{999999, 1000000, 1000000, -1, -1},
		},
	} {
		result := NetworkPackProc(tst.arrival, tst.duration, tst.bufSize)
		require.Equal(t, tst.expected, result)
	}
}

func TestNetworkPackProc6(t *testing.T) {
	for _, tst := range [...]test{
		{
			arrival:  []int{999999, 1000000, 1000000, 1000000, 1000000},
			duration: []int{1, 0, 1, 0, 0},
			bufSize:  1,
			expected: []int{999999, 1000000, 1000000, -1, -1},
		},
	} {
		result := NetworkPackProc(tst.arrival, tst.duration, tst.bufSize)
		require.Equal(t, tst.expected, result)
	}
}
