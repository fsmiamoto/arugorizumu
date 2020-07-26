package stairs

func climbStairsKSteps(n int, k int) int {
	dp := make([]int, k)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j < k; j++ {
			if i-j < 0 {
				continue
			}
			dp[i%k] += dp[(i-j)%k]
		}
	}
	return dp[n%k]
}
