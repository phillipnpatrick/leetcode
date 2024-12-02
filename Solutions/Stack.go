package Solutions

import (
	"fmt"
	"errors"
	"strings"
)

// Stack is a generic stack data structure
type Stack[T any] struct {
	items []T
}

// Push adds an element to the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element from the stack
func (s *Stack[T]) Pop() (T, error) {
	var zeroValue T
	if len(s.items) == 0 {
		return zeroValue, errors.New("stack is empty")
	}

	// Get the last element and remove it from the stack
	topElement := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return topElement, nil
}

// Peek returns the top element without removing it from the stack
func (s *Stack[T]) Peek() (T, error) {
	var zeroValue T
	if len(s.items) == 0 {
		return zeroValue, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// String returns a string representation of the stack
func (s *Stack[T]) String() string {
	var sb strings.Builder
	// sb.WriteString("Stack: [")
	// for i, item := range s.items {
		// if i > 0 {
		// 	sb.WriteString(", ")
		// }
	// 	sb.WriteString(fmt.Sprintf("%v", item))
	// }
	// sb.WriteString("]")

	for _, item := range s.items {
		sb.WriteString(fmt.Sprintf("%v", item))
	}
	return sb.String()
}