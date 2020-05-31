package merge

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

const SliceLength = 500000

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

func BenchmarkSort(b *testing.B) {
	slice := rand.Perm(SliceLength)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sort(slice)
	}
}
