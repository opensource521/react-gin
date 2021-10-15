package engine

type Stack []*Node

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
  return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(node *Node) {
  *s = append(*s, node) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (*Node, bool) {
  if s.IsEmpty() {
    return nil, false
  } else {
    index := len(*s) - 1 // Get the index of the top most element.
    element := (*s)[index] // Index into the slice and obtain the element.
    *s = (*s)[:index] // Remove it from the stack by slicing it off.
    return element, true
  }
}

// Get element at the top of the stack
func (s *Stack) Head() (*Node) {
  if s.IsEmpty() {
    return nil
  } else {
    return (*s)[len(*s) - 1]
  }
}
