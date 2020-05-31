package quick

import (
	"math/rand"
)

func SortRandomized(nums []int) {
	quicksortRandomized(nums, 0, len(nums)-1)
}

func Sort(nums []int) {
	quicksort(nums, 0, len(nums)-1)
}

func quicksortRandomized(nums []int, start, end int) {
	if start < end {
		p := partitionRandomized(nums, start, end)
		quicksortRandomized(nums, start, p-1)
		quicksortRandomized(nums, p+1, end)
	}
}

func quicksort(nums []int, start, end int) {
	if start < end {
		p := partition(nums, start, end)
		quicksortRandomized(nums, start, p-1)
		quicksortRandomized(nums, p+1, end)
	}
}

func partitionRandomized(nums []int, start, end int) int {
	randomIndex := rand.Intn(end-start+1) + start
	nums[randomIndex], nums[end] = nums[end], nums[randomIndex]

	return partition(nums, start, end)
}

func partition(nums []int, start, end int) int {
	pivot := nums[end]

	j := start
	for i := start; i < end; i++ {
		if nums[i] <= pivot {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}

	nums[j], nums[end] = nums[end], nums[j]
	return j
}
