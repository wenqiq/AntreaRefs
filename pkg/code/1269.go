package code

func numWays(steps int, arrLen int) int {
	MOD := 1000000007
	maxLen := min(steps/2, arrLen)

	dp := make([][]int, steps+1)
	for i := 0; i < steps+1; i++ {
		dp[i] = make([]int, maxLen+1)
	}
	dp[0][0] = 1

	for i := 1; i < steps+1; i++ {
		for j := 0; j < maxLen+1; j++ {
			dp[i][j] = dp[i-1][j]
			if j > 0 {
				dp[i][j] += dp[i-1][j-1]
			}
			if j < maxLen-1 {
				dp[i][j] += dp[i-1][j+1]
			}
			dp[i][j] %= MOD
		}
	}
	// fmt.Println(dp)
	return dp[steps][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
