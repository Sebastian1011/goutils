package container

import "errors"

type Stack []interface{}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Top() interface{} {
	if s.Len() == 0 {
		return nil
	}
	a := *s
	return a[s.Len() - 1]
}

func (s *Stack) Push(x interface{}) {
	*s = append(*s, x)
}

func (s *Stack) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

type IntStack []int

func (s *IntStack) Len() int {
	return len(*s)
}

func (s *IntStack) Top() int {
	if s.Len() == 0 {
		panic(errors.New("stack is empty"))
	}
	a := *s
	return a[s.Len() - 1]
}
func (s *IntStack) Push(x int) {
	*s = append(*s, x)
}

func (s *IntStack) Pop() int {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}
