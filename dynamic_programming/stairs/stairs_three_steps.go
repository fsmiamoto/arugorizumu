package stairs

/*
Consider a stair with i steps
f(i) is the number of distinct ways to get to the i-th stair
You can either jump 1 or 2 or 3 stairs at a time.

We want to calculate f(n)

Framework for solving DP problems:
	1. Define the objective function
		f(i) is the number of distinct ways to reach the i-th stair.
	2. Identify the base cases
		f(0) = 1
		f(1) = 1
		f(2) = 22
	3. Write down a recurrence relation for the optimized objective function
		f(n) = f(n-1) + f(n-2) + f(n-3)
	4. What's the order of execution?
		bottom-up
	5. Where to look for the answer?
		f(n)
*/

//  Time complexity: O(n)
//  Space complexity: O(1)
func climbStairsThreeSteps(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	var dp [3]int

	dp[0] = 1
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[2], dp[1], dp[0] = dp[2]+dp[1]+dp[0], dp[2], dp[1]
	}

	return dp[2]
}
