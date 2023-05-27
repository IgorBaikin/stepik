package maxStack

type Node struct {
	Value       int
	CurMaxValue int
	Next        *Node
}

type MaxStack struct {
	stack  *Node
	Length int
}

func (s *MaxStack) Push(elem int) {
	if s.stack != nil {
		temp := &Node{Value: elem, CurMaxValue: max(elem, s.stack.CurMaxValue), Next: s.stack}
		s.stack = temp
	} else {
		s.stack = &Node{Value: elem, CurMaxValue: elem}
	}

	s.Length++
}

func (s *MaxStack) Pop() {
	if s.stack != nil {
		s.stack = s.stack.Next
		s.Length--
	}
}

func (s *MaxStack) GetMax() int {
	if s.stack != nil {
		return s.stack.CurMaxValue
	}

	return 0
}

func (s *MaxStack) Len() int {
	return s.Length
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
