package merge

func Sort(nums []int) {
	mergesort(nums, 0, len(nums)-1)
}

func mergesort(nums []int, start, end int) {
	if start < end {
		middle := (start + end) / 2

		mergesort(nums, start, middle)
		mergesort(nums, middle+1, end)

		merge(nums, start, middle, end)
	}
}

func merge(nums []int, left, middle, right int) {
	lenLeft := middle - left + 1
	lenRight := right - middle

	// Make a copy of left and right halves
	leftHalf := make([]int, lenLeft)
	rightHalf := make([]int, lenRight)

	for i := range leftHalf {
		leftHalf[i] = nums[left+i]
	}

	for j := range rightHalf {
		rightHalf[j] = nums[middle+1+j]
	}

	i, j := 0, 0
	k := left
	for i < lenLeft && j < lenRight {
		if leftHalf[i] <= rightHalf[j] {
			nums[k] = leftHalf[i]
			i++
		} else {
			nums[k] = rightHalf[j]
			j++
		}
		k++
	}

	for i < lenLeft {
		nums[k] = leftHalf[i]
		i++
		k++
	}

	for j < lenRight {
		nums[k] = rightHalf[j]
		j++
		k++
	}
}
