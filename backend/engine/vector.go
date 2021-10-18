package engine

type Vector []int

// Push element into vector
func (v *Vector) Push(num int) {
	*v = append(*v, num)
}

// Check if vector has a certain element
func (v *Vector) Has(num int) bool {
	for _, x := range *v {
		if x == num {
			return true
		}
	}

	return false
}