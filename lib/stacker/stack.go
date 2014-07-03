package stack

import "errors"

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("No top")
	}
	return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	ourStack := *stack
	if len(ourStack) == 0 {
		return nil, errors.New("No Pop")
	}
	x := ourStack[len(ourStack)-1]
	*stack = ourStack[:len(ourStack)-1]
	return x, nil
}
