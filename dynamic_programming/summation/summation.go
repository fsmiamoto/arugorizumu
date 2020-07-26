package main

import "fmt"

// Suppose we don't know how to calculate an arithmetic progression,
// and we would like to solve it with DP
// f(n) = 1 + 2 + 3 + ... + n
// Recurrence relation:
// f(n) = f(n-1) + n

// Of course, this isn't the best solution to this problem.

func nSum(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + i
	}

	return dp[n]
}

func main() {
	fmt.Println("vim-go")
}
