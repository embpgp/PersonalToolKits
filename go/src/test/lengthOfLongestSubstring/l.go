package l

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	m := make(map[byte]int)
	max, l, r := 0, 0, 0
	for ; r < n; r++ {
		if pos, ok := m[s[r]]; ok && pos >= l {
			l = pos + 1
		}
		m[s[r]] = r
		if r-l+1 > max {
			max = r - l + 1
		}
	}

	return max

}

func lengthOfLongestSubstringv2(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	m := make(map[byte]struct{})
	max, l, r := 0, 0, -1
	for r+1 < n {
		if _, ok := m[s[r+1]]; ok {
			delete(m, s[l])
			l++
		} else {
			m[s[r+1]] = struct{}{}
			r++
		}

		if r-l+1 > max {
			max = r - l + 1
		}
	}

	return max

}
