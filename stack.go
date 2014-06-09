package bf_interpreter

type Stack struct {
	stack_ []int
}

func (s Stack) top() int {
	return s.stack_[len(s.stack_)]
}

func (s Stack) push(item int) {
	s.stack_ = append(s.stack_, item)
}

func (s Stack) pop() {
	s.stack_ = s.stack_[:len(s.stack_)-1]
}

func NewStack() *Stack {
	return new(Stack)
}
