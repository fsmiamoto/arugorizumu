package quick

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

const SliceLength = 500000

func TestSortRandomized(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		slice := rand.Perm(SliceLength)
		SortRandomized(slice)
		if !sort.IntsAreSorted(slice) {
			t.Errorf("Expected the slice to be sorted")
		}
	}
}

func TestSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		slice := rand.Perm(SliceLength)
		Sort(slice)
		if !sort.IntsAreSorted(slice) {
			t.Errorf("Expected the slice to be sorted")
		}
	}
}

func BenchmarkSortRandomized(b *testing.B) {
	slice := rand.Perm(SliceLength)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortRandomized(slice)
	}
}

func BenchmarkSort(b *testing.B) {
	slice := rand.Perm(SliceLength)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sort(slice)
	}
}
