package merge

import (
	"sort"
	"testing"

	s "github.com/fsmiamoto/arugorizumu/sort"
)

func TestSort(t *testing.T) {
	for i := 0; i < 10; i++ {
		slice := s.RandomInts(s.SliceLength, s.MaxInt)
		Sort(slice)
		if !sort.IntsAreSorted(slice) {
			t.Errorf("Expected the slice to be sorted")
		}
	}
}

func BenchmarkSort(b *testing.B) {
	slice := s.RandomInts(s.SliceLength, s.MaxInt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sort(slice)
	}
}
