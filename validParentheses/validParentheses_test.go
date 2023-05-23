package validParentheses

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	input    string
	expected string
}

func TestValidParentless(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    "([](){([])})",
			expected: "Success",
		},
		{
			input:    "()[]}",
			expected: "5",
		},
		{
			input:    "{{[()]]",
			expected: "7",
		},
		{
			input:    "{{{[][][]",
			expected: "3",
		},
		{
			input:    "{*{{}",
			expected: "3",
		},
		{
			input:    "[[*",
			expected: "2",
		},
		{
			input:    "{*}",
			expected: "Success",
		},
		{
			input:    "{{",
			expected: "2",
		},
		{
			input:    "{}",
			expected: "Success",
		},
		{
			input:    "",
			expected: "Success",
		},
		{
			input:    "}",
			expected: "1",
		},
		{
			input:    "*{}",
			expected: "Success",
		},
		{
			input:    "{{{**[][][]",
			expected: "3",
		},
		{
			input:    "()({}",
			expected: "3",
		},
		{
			input:    "{{[()]}",
			expected: "1",
		},
		{
			input:    "[]",
			expected: "Success",
		},
		{
			input:    "{}[]",
			expected: "Success",
		},
		{
			input:    "[()]",
			expected: "Success",
		},
		{
			input:    "(())",
			expected: "Success",
		},
		{
			input:    "{[]}()",
			expected: "Success",
		},
		{
			input:    "([](){([])})",
			expected: "Success",
		},
		{
			input:    "foo(bar);",
			expected: "Success",
		},
		{
			input:    "{",
			expected: "1",
		},
		{
			input:    "{[}",
			expected: "3",
		},
		{
			input:    "()[]}",
			expected: "5",
		},
		{
			input:    "{{[()]]",
			expected: "7",
		},
		{
			input:    "foo(bar[i);",
			expected: "10",
		},
		{
			input:    "[]([]",
			expected: "3",
		},
	} {
		result := ValidParentless(tst.input)
		require.Equal(t, tst.expected, result)
	}
}
