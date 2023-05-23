package validParentheses

import (
	"fmt"
)

type Stack struct {
	Index int
	Ch    rune
}

func ValidParentless(s string) string {
	m := make(map[rune]rune)
	m['('] = ')'
	m['{'] = '}'
	m['['] = ']'
	stack := make([]Stack, 0)
	rs := []rune(s)
	var i int
	for i = 0; i < len(s); i++ {
		if _, ok := m[rs[i]]; !ok && rs[i] != ')' && rs[i] != '}' && rs[i] != ']' {
			continue
		}

		if value, ok := m[rs[i]]; ok {
			stack = append(stack, Stack{Index: i, Ch: value})
			continue
		}

		if l := len(stack); l == 0 || stack[l-1].Ch != rs[i] {
			return fmt.Sprint(i + 1)
		}

		stack = stack[:len(stack)-1]

	}

	if len(stack) == 0 {
		return "Success"
	} else {
		return fmt.Sprint(stack[len(stack)-1].Index + 1)
	}
}
