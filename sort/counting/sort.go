package counting

import (
	s "github.com/fsmiamoto/arugorizumu/sort"
)

// Given the nature of counting algorithms, we need to define a maximum
// integer value
const MaxInt = s.MaxInt

// Counting Sort consumes len(elems) + MaxInt units of time
// Given a 'small' MaxInt, Counting Sort achieves sorting in linear time.
func Sort(elems []int) {
	countingSort(elems, MaxInt)
}

// Sort an slice of integers where each element
// is in the interval 0..max-1
func countingSort(elems []int, max int) {
	frequencies := make([]int, max+1)
	aux := make([]int, len(elems))

	for i := range elems {
		frequencies[elems[i]+1]++
	}
	// frequencies[r] is now the frequencies of r-1 in elems

	for r := 1; r <= max; r++ {
		frequencies[r] += frequencies[r-1]
	}
	// frequencies[r] is now the frequencies of predecessors of r
	// Therefore, elements equal to r should begin at index frequencies[r]

	for i := range elems {
		aux[frequencies[elems[i]]] = elems[i]
		// This sets the index for the next element with the same value
		frequencies[elems[i]]++
	}

	for i := range elems {
		elems[i] = aux[i]
	}
}
