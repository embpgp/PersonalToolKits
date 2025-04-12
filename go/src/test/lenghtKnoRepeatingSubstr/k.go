package lenghtknorepeatingsubstr

func numKLenSubstrNoRepeats(s string, k int) int {
	//假设都是小写字母
	for k > 26 {
		return 0
	}
	ans := 0
	n := len(s)
	freq := make([]int, 26) //记录每个字母出现的次数
	left, right := 0, 0
	for right < n {
		freq[s[right]-'a']++
		for freq[s[right]-'a'] > 1 {
			freq[s[left]-'a']--
			left++
		}
		if right-left+1 == k {
			ans++
			freq[s[left]-'a']--
			left++
		}
		right++
	}
	return ans
}

// 求长度为K的无重复子串的个数
func f(s string, k int) int {

	if k > 26 || k > len(s) {
		return 0
	}
	ans := 0
	left, right := 0, 0
	freq := make([]byte, 26)
	for right < len(s) {
		freq[s[right]-'a']++
		for freq[s[right]-'a'] > 1 {
			//left右边走
			freq[s[left]-'a']--
			left++
		}
		if right-left+1 == k {
			ans++
			freq[s[left]-'a']--
			left++
		}
		right++
	}
	return ans
}
