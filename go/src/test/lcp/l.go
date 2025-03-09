package l

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	///横向比较每个字符串
	max := strs[0]
	for i := 1; i < len(strs); i++ {
		max = lcp(max, strs[i])
		if len(max) == 0 {
			break
		}
	}
	return max

}

func lcp(s1, s2 string) string {
	n := len(s1)
	if len(s2) < n {
		n = len(s2)
	}
	ret := s1[:n]
	for i := 0; i < n; i++ {
		if s1[i] == s2[i] {
			continue
		}
		return s1[:i]
	}
	return ret

}

func longestCommonPrefixV2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	// 纵向比较
	s1 := strs[0]
	for i := 0; i < len(s1); i++ {
		for j := 1; j < len(strs); j++ {
			if i+1 > len(strs[j]) || s1[i] != strs[j][i] {
				return s1[:i]
			}
		}
	}
	return s1

}
