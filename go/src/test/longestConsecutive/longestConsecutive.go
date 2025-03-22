package longestConsecutive

func longestConsecutive(nums []int) int {
	n := len(nums)
	if n == 0 || n == 1 {
		return n
	}
	nMap := make(map[int]bool, n)
	for _, v := range nums {
		nMap[v] = true
	}
	//遍历map，理论是当i-1不存在的时候，开始统计
	ans := 0
	for i := range nMap {
		if !nMap[i-1] {
			curr := 1
			idx := i
			for nMap[idx+1] {
				curr++
				idx++
			}
			if curr > ans {
				ans = curr
			}
		}
	}
	return ans
}
