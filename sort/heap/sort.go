package heap

import (
	h "github.com/fsmiamoto/heap"
)

func Sort(nums []int) {
	hp := h.New(nil, len(nums), h.MinInt)

	for i := range nums {
		hp.Insert(nums[i])
	}

	for i := range nums {
		v, _ := hp.Extract()
		nums[i] = v.(int)
	}
}
