package main

type LPS struct {
	S string
}

// 动态规划
func (m *LPS) dp() string {
	if len(m.S) < 2 {
		return m.S
	}
	dp := make([][]int, len(m.S))
	for i := 0; i < len(m.S); i++ {
		dp[i] = make([]int, len(m.S))
	}

	for i := 0; i < len(m.S); i++ {
		dp[i][i] = 1
	}
	maxLen := 1
	begin := 0
	for L := 2; L <= len(m.S); L++ { //列
		for i := 0; i < len(m.S); i++ { //行，遍历右上角，每次是上到下，从左到右
			j := i + L - 1
			if j >= len(m.S) {
				break
			}
			if m.S[i] != m.S[j] {
				dp[i][j] = 0
			} else {
				if j-i < 3 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] == 1 && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}

	}

	return m.S[begin : begin+maxLen]
}

// 中心扩展
func (m *LPS) center() string {
	if len(m.S) < 2 {
		return m.S
	}
	start, end := 0, 0
	for i := 0; i < len(m.S); i++ {
		s1, e1 := calcCenterMaxLen(m.S, i, i)
		s2, e2 := calcCenterMaxLen(m.S, i, i+1)
		if e1-s1 > end-start {
			start, end = s1, e1
		}
		if e2-s2 > end-start {
			start, end = s2, e2
		}
	}
	return m.S[start : end+1]
}

func calcCenterMaxLen(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
