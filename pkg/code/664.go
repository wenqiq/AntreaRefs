package code

import (
	"math"
)

func strangePrinter(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			dp[i][j] = math.MaxInt32
			dp[j][j] = 1
			if s[i] == s[j] {
				dp[i][j] = dp[i][j-1]
			} else {
				for k := i; k < j; k++ {
					dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
				}
			}
		}
	}

	return dp[0][n-1]
}
