package engine

type Queue []int

// IsEmpty: check if queue is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// Push a new value onto the queue
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Remove and return first element of queue. Return false if queue is empty.
func (s *Queue) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	} else {
		element := (*s)[0]
		*s = (*s)[1:]
		return element, true
	}
}

// Get the first element of queue
func (s *Queue) Front() int {
	if s.IsEmpty() {
		return -1
	} else {
		return (*s)[0]
	}
}
