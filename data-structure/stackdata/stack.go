package stackdata

import "fmt"

/// Stack Data Structure in Golang
// Stack Data Structure

type node[T any] struct {
	data T
	next *node[T]
}

func newNode[T any](data T) *node[T] {
	return &node[T]{
		data: data,
		next: nil,
	}
}

type Stack[T any] struct {
	top  *node[T]
	size uint
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		top:  nil,
		size: 0,
	}
}

// Public function of stack

// Size get the stack length
func (s *Stack[T]) Size() uint {
	return s.size
}

// Push push data to the top of the stack
func (s *Stack[T]) Push(data T) *Stack[T] {
	newNode := newNode[T](data)
	newNode.next = s.top
	s.top = newNode
	s.size += 1
	return s
}

// Pop pop the top node of the stack
func (s *Stack[T]) Pop() T {
	result := s.Peek()
	s.top = s.top.next
	s.size -= 1
	return result
}

// IsEmpty return true if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.size == 0
}

// Peek return the top node value
func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue
	}
	return s.top.data
}

// PrintStack print the stack
func (s *Stack[T]) PrintStack() {
	curr := s.top
	fmt.Print("Stack: ")
	for curr != nil {
		fmt.Print(curr.data, "->")
		curr = curr.next
	}
	fmt.Print("nil\n")
}
