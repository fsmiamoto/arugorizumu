package stairs_with_cost

/*
Problem:
	Paid Staircase

	You are climbing a paid staircase. It takes n steps to reach to the top and you have
	to pay p[i] to step on the i-th stair. Each time you can climb 1 or k steps.
	What's the cheapest amount you ahve to pay to get to the top of the staircase?

	Example: n = 3, k = 2, price = [0,3,2,4] -> $6  = $2 + $4 (Step on 2 and then step on 3)
*/

// Time complexity: O(nk)
// Space complexity: O(n), but it could be O(k)
func costOfClimbingToStair(n int, k int, price []int) int {
	memo := make([]int, n+1)
	memo[0] = 0

	for i := 1; i <= n; i++ {
		if i <= k {
			memo[i] = price[i]
		} else {
			memo[i] = minOf(memo[i-k:i]) + price[i]
		}
	}

	return memo[n]
}

// Time complexity O(nk)
// Space complexity: O(k)
func pathOfClimbingToStair(n int, k int, price []int) []int {
	memo := make([]int, 0, k)

	from := n
	for from > 0 {
		memo = append(memo, from)
		if from-k < 0 {
			from = indexOfMin(price[0:from])
			continue
		}
		// Offset of from-k
		from = indexOfMin(price[from-k:from]) + from - k
	}

	memo = append(memo, 0)
	reverse(memo)

	return memo
}

// minOf deliberately doesn't have safety checks
func minOf(s []int) int {
	min := s[0]
	for i := range s {
		if s[i] < min {
			min = s[i]
		}
	}
	return min
}

func indexOfMin(s []int) int {
	min := 0
	for i := range s {
		if s[i] < s[min] {
			min = i
		}
	}
	return min
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
