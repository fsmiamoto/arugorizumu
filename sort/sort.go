package sort

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const SliceLength = 500000
const MaxInt = 1000

func RandomInts(size int, maxValue int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(maxValue)
	}
	return slice
}
