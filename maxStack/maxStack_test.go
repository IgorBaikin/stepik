package maxStack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxStack(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		stack := MaxStack{}

		require.Equal(t, 0, stack.Len())
		require.Equal(t, 0, stack.GetMax())
	})

	t.Run("push to empty stack", func(t *testing.T) {
		stack := MaxStack{}
		stack.Push(10)
		require.Equal(t, 10, stack.GetMax())
	})

	t.Run("remove items", func(t *testing.T) {
		stack := MaxStack{}

		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		stack.Push(4)

		require.Equal(t, 4, stack.GetMax())
		stack.Pop()
		require.Equal(t, 3, stack.GetMax())

		stack.Pop()
		stack.Pop()
		stack.Pop()
		stack.Push(5)
		require.Equal(t, 5, stack.GetMax())
	})

	t.Run("remove items", func(t *testing.T) {
		stack := MaxStack{}

		stack.Push(4)
		stack.Push(3)
		stack.Push(2)
		stack.Push(1)

		require.Equal(t, 4, stack.GetMax())
		stack.Pop()
		require.Equal(t, 4, stack.GetMax())

		stack.Pop()
		stack.Pop()
		stack.Pop()
		stack.Push(5)
		require.Equal(t, 5, stack.GetMax())
	})
}
