package longestrepeatingsubstring

func longestRepeatingSubstring(s string) int {
	// 动态规划，定义dp[i][j]为分别为i结尾和j结尾的时候，
	// 当前重复子串，因此若s[i]=s[j],则dp[i][j]=d[i-1][j-1]+1
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 0
		}
	}
	ret := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				if i > 0 { //处理特殊case，比如aaaaa
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = 1
				}

			}
			if dp[i][j] > ret {
				ret = dp[i][j]
			}
		}
	}
	return ret
}
